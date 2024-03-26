package gapi

import (
	"fmt"

	db "go.mod/db/sqlc"
	"go.mod/pb"
	"go.mod/token"
	"go.mod/utils"
	"go.mod/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleprojectServer
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	taskDistributor worker.TaskDistributor
}


func NewServer(config utils.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
