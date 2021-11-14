package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("first try to use RabbitMQ")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(1)
	}
	// defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()
	q, err := ch.QueueDeclare(
		"Test Queue",
		false,
		false,
		false,
		false,
		nil,
	)

	fmt.Println(q)
	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",
		"Test Queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully published message to queue")
}
