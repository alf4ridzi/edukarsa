package main

import (
	"edukarsa-backend/api/routes"
	"edukarsa-backend/internal/config"
	"edukarsa-backend/pkg/postgresql"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgresql.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := postgresql.Migrate(db); err != nil {
		log.Fatal(err)
	}

	routes.Setup(cfg, db)
}
