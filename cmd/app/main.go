package main

import (
	"future-path/internal/handler/rest"
	"future-path/internal/repository"
	"future-path/internal/service"
	"future-path/pkg/config"
	"future-path/pkg/database/mariadb"
	"log"
)

func main() {
	config.LoadEnvironment()
	db, err := mariadb.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}

	if err := mariadb.Migrate(db); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	r := rest.NewRest(svc)
	r.MountEndpoint()
	r.Run()
}
