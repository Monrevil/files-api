package main

import (
	"github.com/streadway/amqp"
)


//RabbitMQ channels are not thread-safe, connections are
//Each hit of an endpoint should open new channel and Queue
func initRabbit(path string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"imageID", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	rabbitPublishText(ch, q, path)
}

func rabbitPublishText(ch *amqp.Channel, q amqp.Queue, id string) {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(id),
		})
	failOnError(err, "Failed to publish a message")
}
