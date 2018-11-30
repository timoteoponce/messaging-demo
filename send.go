package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"brian", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter message body: ")
	body, _ := reader.ReadString('\n')
	for i := 0; i < 1000; i++ {
		finalMsg := fmt.Sprintf("[%d]%s", i, body)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(finalMsg),
			})
		log.Printf(" [x] Sent %s %s", i, finalMsg)
		failOnError(err, "Failed to publish a message")
	}
}
