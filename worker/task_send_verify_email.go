package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db "go.mod/db/sqlc"
	"go.mod/utils"
)

const TaskSenderVerfiyEmail = "task:send_verify_email" // important to asynq to recoginze what kind of task it is distributing or processing

// this function conatiin all the information that we want to store in the rredis
type PayloadSendVerifyEmail struct {
	UserName string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerfiyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option) error {


	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf(("can't marshal"))
	}
	task := asynq.NewTask(TaskSenderVerfiyEmail, jsonPayload, opts...) // to create new task
	//send this task to redis queue
	info, err := distributor.client.EnqueueContext(ctx, task)

	if err != nil {
		return fmt.Errorf("Failed to enque task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("Queue", info.Queue).Int("Max_retry", info.MaxRetry).Msg("enqued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error{
	var payload PayloadSendVerifyEmail
	if err :=json.Unmarshal(task.Payload(),&payload); err!=nil{
		return fmt.Errorf("failed to unmarshl payload %w", asynq.SkipRetry)
	}
	user,err:=processor.store.GetUser(ctx, payload.UserName)
		if err != nil{
			if err == sql.ErrNoRows{
				return fmt.Errorf("user dosen't exist: %w", asynq.SkipRetry)
			}
			return fmt.Errorf("failed to get user: %w", err)
		}

		verifyEmail,err := processor.store.CreateVerifyEmail(ctx,db.CreateVerifyEmailParams{
			Username:user.Username,
			Email:user.Email,
			SecretCode:utils.RandomString(32),
		})
		if err != nil{
			return fmt.Errorf("Failed to verify emails: %w",err)
		}
		subject := "Welcome to Simple Bank"
		// replace this URL with an environment variable that points to a front-end page
		verifyUrl := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s",
			verifyEmail.ID, verifyEmail.SecretCode)
		content := fmt.Sprintf(`Hello %s,<br/>
		Thank you for registering with us!<br/>
		Please <a href="%s">click here</a> to verify your email address.<br/>
		`, user.FullName, verifyUrl)
		to := []string{user.Email}
	
		err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to send verify email: %w", err)
		}
		
		log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("processd task")
	  return nil
 }
