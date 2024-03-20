package gapi

import (
	"context"
	"errors"

	db "go.mod/db/sqlc"
	"go.mod/pb"
	"go.mod/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) Loginuser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "Notfound: %s", err)
	
		}
		return nil, status.Errorf(codes.Internal, "failed to find user: %s", err)
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect pssword: %s", err)
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create acess token: %s", err)
	}
   
	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to creste refresh user: %s", err)
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}
	rsp := &pb.LoginUserResponse{
		User: convertUser(user),
		SessionId: session.Username,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
	}
	return rsp,nil
}