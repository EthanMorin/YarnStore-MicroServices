package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	Register()
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Println(err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err.Error())
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"order", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Println(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Println(err.Error())
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Sending email confirmation to: %s", d.Body)
			// send email
		}
	}()

	log.Print("Waiting for messages")
	<-forever
}