package main

import (
	"context"
	"database/sql"
	"net"

	"log"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	db "go.mod/db/sqlc"
	"go.mod/gapi"
	"go.mod/pb"
	"go.mod/utils"
	"google.golang.org/grpc"
)

 var testqueries *db.Queries
 
 func main(){
    config, err := utils.LoadConfig("../..")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

	conn,err:=sql.Open(config.DBDriver,config.DBSource)
    // connConfig, err := pgx.ParseConfig(config.DBSource)
    // if err != nil {
    //     log.Fatal("unable to parse config:", err)
    // }
   // conn, err := pgx.ConnectConfig(context.Background(), connConfig)
    if err != nil {
        log.Fatal("unable to connect:", err)
    }

   // testqueries = db.New(conn)
	store:=db.NewStore(conn)
	runGrpcServer(config,store)
    


}

func runGrpcServer(config utils.Config, store db.Store){

server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	grpcserver := grpc.NewServer()
	pb.RegisterSimpleprojectServer(grpcserver,server)


	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}
	err = grpcserver.Serve(listener)
	
	if err != nil {
		log.Fatal("cannot start gRbc server")
	}
}
