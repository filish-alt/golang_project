package gapi

import (
	"context"

	db "go.mod/db/sqlc"
	"go.mod/pb"
	val "go.mod/validation"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	violations := validateEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	
	txResult,err := server.store.VerifyEmailTx(ctx,db.VerifyEmailTxParams{
		EmailId: req.GetEmailId(),
		SecretCode: req.GetSecretCode(),
	})
	if err != nil{
		return nil, status.Errorf(codes.Internal,"failed to verify email")
	}
	
	rsp := &pb.VerifyEmailResponse{
		IsVerified: txResult.User.IsEmailVerified,
		
	}
	return rsp,nil
}


func validateEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, fieldViolation("emailId", err))
	}

	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err))
	}

	// if err := val.ValidateEmail(req.GetEma); err != nil {
	// 	violations = append(violations, fieldViolation("Email", err))
	// }

	

	return violations
}






