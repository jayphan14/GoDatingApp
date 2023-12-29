package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/datingdb?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("DB can not be connected", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
