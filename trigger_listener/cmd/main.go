package main

import (
	"rabbitmq/trigger_listener/internal/repository"
	"rabbitmq/trigger_listener/internal/usecase"
)

func main() {
	repo := repository.NewRestRepository()
	trigger := usecase.NewTriggerUsecase(repo)
	trigger.RunTrigger("amqp://guest:guest@localhost:5672/")
}
