package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to our RabbitMQ instance")

	channel, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}

	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(queue)

	err = channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Successfully publish message to queue")
}
