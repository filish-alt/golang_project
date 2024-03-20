package gapi

import (
	"context"

	db "go.mod/db/sqlc"
	"go.mod/pb"
	"go.mod/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server Server) Createuser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed hash password: %s", err)
	
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	rsp :=&pb.CreateUserResponse{
		User: convertUser(user),
	}
	return rsp,nil
}

func Get(s string) {
	panic("unimplemented")
}





