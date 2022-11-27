package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"rabbitmq/trigger_listener/internal/entity/global"
	"rabbitmq/trigger_listener/internal/repository"

	"github.com/streadway/amqp"
)

type TriggerUsecase struct {
	restRepo repository.RestRepository
}

func NewTriggerUsecase(restRepo repository.RestRepository) TriggerUsecase {
	return TriggerUsecase{
		restRepo: restRepo,
	}
}

func (u *TriggerUsecase) RunTrigger(connString string) {

	conn, err := amqp.Dial(connString) // Создаем подключение к RabbitMQ
	if err != nil {
		log.Fatalf("unable to open connect to RabbitMQ server. Error: %s", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"trigger",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	var requestID int
	go func() {
		for d := range msgs {

			if err := json.Unmarshal(d.Body, &requestID); err != nil {
				log.Fatalln(err)
			}

			if err = ch.Publish(
				global.ExchangeName,
				global.KeyDebug,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        d.Body,
				},
			); err != nil {
				log.Fatalln(err)
			}

			data, err := u.restRepo.FindPhoneNumber(requestID)
			switch err {
			case nil:
				body, err := json.Marshal(data)
				if err != nil {
					log.Fatalln(err)
				}
				if err = ch.Publish(
					global.ExchangeName,
					global.KeyInfo,
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        body,
					},
				); err != nil {
					log.Fatalln(err)
				}

			case global.ErrNoData:
				if err = ch.Publish(
					global.ExchangeName,
					global.KeyError,
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(fmt.Sprintf("error: Not found %d", requestID)),
					},
				); err != nil {
					log.Fatalln(err)
				}

			case global.ErrInternalError:
				if err = ch.Publish(
					"",
					"trigger",
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(fmt.Sprintf("%d", requestID)),
					},
				); err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()

	<-forever
}
