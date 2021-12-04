package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")
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

	messages, err := channel.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("Received Message : %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to out RabbitMQ Instance")
	fmt.Println(" [*] - waiting for message ")
	<-forever

}
