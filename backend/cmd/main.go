package main

import (
	"edukarsa-backend/api/routes"
	"edukarsa-backend/internal/config"
	"edukarsa-backend/pkg/postgresql"
	"edukarsa-backend/pkg/postgresql/seeders"
	"flag"
	"log"
)

func main() {
	seedFlag := flag.Bool("seed", false, "run database seed")
	flag.Parse()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	config.AppConfig = cfg

	db, err := postgresql.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer postgresql.Close(db)

	if err := postgresql.Migrate(db); err != nil {
		log.Fatal(err)
	}

	if *seedFlag {
		log.Println("running database seed...")
		seeder := seeders.Seed{DB: db}

		if err := seeder.Run(); err != nil {
			log.Fatal(err)
		}
		return
	}

	routes.SetupRoute(cfg, db)
}
