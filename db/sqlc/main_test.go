package db

import (

	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"go.mod/utils"
)

 var testqueries *Queries
 
 func TestingMain(m *testing.M) {
    config, err := utils.LoadConfig("../..")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }


    testqueries = New(config)
	os.Exit(m.Run())
    


}