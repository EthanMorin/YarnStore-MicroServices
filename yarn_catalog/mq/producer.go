package services

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func producer() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalln(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalln(err)
	}

	q, err := ch.QueueDeclare(
		"cart",
		false,
		false,
		false,
		false,
		nil,
	)

	log.Println(q)

	if err != nil {
		log.Fatalln(err)
	}
	return ch
}

func PubMsg[T any](obj T) {
	msg, err := json.Marshal(obj)
	if err != nil {
		log.Fatalln(err)
	}
	ch := producer()
	ch.Publish(
		"",
		"cart",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}
