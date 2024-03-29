package gapi

import (
	"context"

	db "go.mod/db/sqlc"
	"go.mod/pb"
	"go.mod/utils"
	val "go.mod/validation"
	"go.mod/worker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server Server) Createuser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed hash password: %s", err)
	
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),
			HashedPassword: hashedPassword,
			FullName:       req.GetFullName(),
			Email:          req.GetEmail(),
		},
	AlterCreate: func(user db.User) error{
		taskpayload := &worker.PayloadSendVerifyEmail{
			UserName:user.Username,
		}
		return server.taskDistributor.DistributeTaskSendVerfiyEmail(ctx,taskpayload)
			
	},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	
	
	rsp :=&pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}
	return rsp,nil
}


func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := val.ValidateFullName(req.GetFullName()); err != nil {
		violations = append(violations, fieldViolation("full_name", err))
	}

	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}

	return violations
}






