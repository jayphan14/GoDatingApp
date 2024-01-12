package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jayphan14/GoDatingApp/util"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, errLoadingConfig := util.LoadConfig("../")
	if errLoadingConfig != nil {
		log.Fatal("cant load config", errLoadingConfig)
	}
	var err error
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("DB can not be connected", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
