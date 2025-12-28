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
	migrationFlag := flag.Bool("migrate", false, "run database migrate")
	dropTableFlag := flag.Bool("wipe", false, "drop all tables")

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

	if *migrationFlag {
		log.Println("running migrate...")
		if err := postgresql.Migrate(db); err != nil {
			log.Fatal(err)
		}

		return
	}

	if *dropTableFlag {
		log.Println("running wipe...")
		if err := postgresql.DropTable(db); err != nil {
			log.Fatal(err)
		}

		return
	}

	enforcer, err := config.InitCasbin(db)
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoute(cfg, db, enforcer)
}
