package db

import (
	"context"

	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"go.mod/utils"
)

 var testqueries *Queries
 
 func TestingMain(m *testing.M) {
    config, err := utils.LoadConfig("../..")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    connConfig, err := pgx.ParseConfig(config.DBSource)
    if err != nil {
        log.Fatal("unable to parse config:", err)
    }
    conn, err := pgx.ConnectConfig(context.Background(), connConfig)
    if err != nil {
        log.Fatal("unable to connect:", err)
    }

    testqueries = New(conn)
	os.Exit(m.Run())
    


}