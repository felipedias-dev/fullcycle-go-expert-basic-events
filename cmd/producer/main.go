package main

import (
	"github.com/felipedias-dev/fullcycle-go-expert-basic-events/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msg := "Hello World"
	amqpMsg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}
	rabbitmq.Publish(ch, amqpMsg, "amq.direct")
}
