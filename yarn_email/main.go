package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"order", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Sending email confirmation to: %s", d.Body)
			SendUserEmail(string(d.Body))
		}
	}()

	log.Print("Waiting for messages")
	<-forever
}

// package main

// import (
// 	"log"

// 	"github.com/streadway/amqp"
// )

// func main() {
// 	Register()
// 	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer conn.Close()

// 	ch, err := conn.Channel()
// 	if err != nil {
// 		log.Println(err.Error())
// 	}
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"order", // name
// 		false,   // durable
// 		false,   // delete when unused
// 		false,   // exclusive
// 		false,   // no-wait
// 		nil,     // arguments
// 	)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	if err != nil {
// 		log.Println(err.Error())
// 	}

// 	var forever chan struct{}

// 	go func() {
// 		for d := range msgs {
// 			log.Printf("Sending email confirmation to: %s", d.Body)
// 			// send email
// 			SendUserEmail(string(d.Body))
// 		}
// 	}()

// 	log.Print("Waiting for messages")
// 	<-forever
// }
