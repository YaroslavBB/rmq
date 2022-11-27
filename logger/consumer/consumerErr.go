package consumer

import (
	"fmt"
	"log"
	"rabbitmq/logger/entity"

	"github.com/streadway/amqp"
)

func StartConsumeErr(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		entity.ExchangeName,
		amqp.ExchangeDirect,
		false,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		entity.KeyError,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	if err = ch.QueueBind(
		q.Name,
		entity.KeyError,
		entity.ExchangeName,
		false,
		nil,
	); err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			fmt.Printf("message [error] -- %s\n", d.Body)
		}
	}()

	return nil
}
