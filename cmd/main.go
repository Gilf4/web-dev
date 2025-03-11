package main

import (
	"GoForBeginner/cmd/api"
	"GoForBeginner/internal/config"
	"GoForBeginner/internal/db"
	"GoForBeginner/internal/db/migrations"
	"golang.org/x/net/context"

	"log"
)

func main() {
	cfg := config.MustLoad()

	pool, err := db.Connect(cfg.DB)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer pool.Close()

	if err := migrations.CreateTables(context.Background(), pool); err != nil {
		log.Fatal(err)
	}

	apiServer := api.NewAPIServer(cfg.Server.Addr, pool)

	if err := apiServer.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
