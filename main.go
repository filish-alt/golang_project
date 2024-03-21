package main

import (
	"context"
	"database/sql"
	"net"
	"net/http"

	"log"

	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	db "go.mod/db/sqlc"
	_ "go.mod/docs/statik"
	"go.mod/gapi"
	"go.mod/pb"
	"go.mod/utils"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	
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

   runDbMigration(config.MigrationURL,config.DBSource)

	store:=*db.NewStore(conn)
	go runGatewayServer(config, store)
	runGrpcServer(config,store)
    


}

func runDbMigration(migrationUrl string, dbsource string){
	migration, err := migrate.New(migrationUrl,dbsource)
	if err != nil{
		log.Fatal("can't migrate new instance", err)
	}
   if err= migration.Up(); err!=nil{
	log.Fatal("error in migration up", err)
   }
   log.Printf("db migrated sucessfully")
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

func runGatewayServer(config utils.Config, store db.Store){

	server, err := gapi.NewServer(config, store)
		if err != nil {
			log.Fatal("cannot create server")
		}

		jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		})
		grpcmux:=runtime.NewServeMux(jsonOption)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
	
	 err = pb.RegisterSimpleprojectHandlerServer(ctx,grpcmux,server)
        
	 if err != nil{
		log.Fatal("can't register handler server")
	 }
	 mux := http.NewServeMux()
//route to the grpcmux
	 mux.Handle("/",grpcmux)

	 statikFS, err := fs.New()
	 if err != nil {
		 log.Fatal("cannot create statik fs")
	 }
 
	 swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	 mux.Handle("/swagger/", swaggerHandler)

		listener, err := net.Listen("tcp", config.HTTPServerAddress)
		if err != nil {
			log.Fatal("cannot create listener")
		}
		err = http.Serve(listener,mux)
		
		if err != nil {
			log.Fatal("cannot start http server")
		}
	}
	
