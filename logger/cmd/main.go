package main

import (
	"log"
	"rabbitmq/logger/consumer"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/") // Создаем подключение к RabbitMQ
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
		panic(err)
	}

	for {
		err = consumer.StartConsumeAll(conn)
		if err != nil {
			log.Fatalln(err)
		}

		err = consumer.StartConsumeErr(conn)
		if err != nil {
			log.Fatalln(err)
		}

		err = consumer.StartConsumeInfo(conn)
		if err != nil {
			log.Fatalln(err)
		}

		err = consumer.StartConsumeDebug(conn)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
