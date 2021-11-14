package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("First try using RabbitMQ")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("failed initializing broker connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"Test Queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	loadForever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println("[*] waiting for messages")
	<-loadForever
}
