package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jayphan14/GoDatingApp/api"
	db "github.com/jayphan14/GoDatingApp/sqlc"
	"github.com/jayphan14/GoDatingApp/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	newServer := api.NewServer(store)
	newServer.Start(config.ServerAddress)
}
