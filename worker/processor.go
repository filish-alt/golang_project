package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	db "go.mod/db/sqlc"
	"go.mod/mail"
)
type TaskProcessor interface{
	Start() error
	Shutdown()
     // this function to process the task send verify email to the interface
	 ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}
type RedisTaskProcessor struct {
	server *asynq.Server
	store db.Store
	mailer mail.EmailSender
}

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)
func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, store db.Store, mailer mail.EmailSender) TaskProcessor{
   server := asynq.NewServer(
	  redisOpt,
	  asynq.Config{
		Queues: map[string]int{
		QueueCritical: 10,
		QueueDefault:  5,
	},//this config object allow to control many different parametr of async server
	ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
		log.Error().Err(err).Str("type", task.Type()).
			Bytes("payload", task.Payload()).Msg("process task failed")
	}),
	Logger:&Logger{},
},

)
   return &RedisTaskProcessor{
	server: server,
	store: store,
	mailer: mailer,
   }
}

func (processor *RedisTaskProcessor) Start() error{
    mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSenderVerfiyEmail,processor.ProcessTaskSendVerifyEmail)
    return processor.server.Start(mux)
}

func (processor *RedisTaskProcessor) Shutdown(){
	processor.server.Shutdown()
}

