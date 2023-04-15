package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, outMsg chan amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,         // queue
		"go-consumer", // consumer
		false,         // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		outMsg <- msg
	}
	return nil
}

func Publish(ch *amqp.Channel, msg amqp.Publishing, exchange string) error {
	return ch.PublishWithContext(
		context.Background(),
		exchange, // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		msg,      // msg
	)
}
