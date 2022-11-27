package main

import (
	"log"
	"rabbitmq/rest_service/config"
	"rabbitmq/rest_service/internal/app/rest"
	"rabbitmq/rest_service/internal/repository/postgres"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.NewConfig("../config/config.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := sqlx.Open("postgres", config.DbConn())
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()

	repo := postgres.NewNumberRepository(db)

	server := rest.NewRest(r, repo)

	server.Run()
}
