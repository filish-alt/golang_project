package gapi

import (
	
	"fmt"
	db "go.mod/db/sqlc"
	"go.mod/pb"
	"go.mod/token"
	"go.mod/utils"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleprojectServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
}

// Createuser implements pb.SimpleprojectServer.
// func (s *Server) Createuser(context.Context, *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
// 	panic("unimplemented")
// }

// Loginuser implements pb.SimpleprojectServer.
// func (s *Server) Loginuser(context.Context, *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
// 	panic("unimplemented")
// }

// mustEmbedUnimplementedSimpleprojectServer implements pb.SimpleprojectServer.
func (s *Server) mustEmbedUnimplementedSimpleprojectServer() {
	panic("unimplemented")
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
