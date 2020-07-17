package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("consuenr application")
	conn, err := amqp.Dial("amqp://admin:password@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("successfully connected to rabbit mq")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"testQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received message : %s\n", d.Body)
		}
	}()

	fmt.Println("successfully connected to our rabbit mq instance")
	fmt.Println("[*] - waiting for message")
	<-forever
}
