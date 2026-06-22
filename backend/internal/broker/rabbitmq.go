package broker

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

var RMQ *RabbitMQ

func InitRabbitMQ(url string) error {
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"chat.logs.queue", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		return err
	}

	RMQ = &RabbitMQ{
		conn:    conn,
		channel: ch,
	}
	log.Println("Connected to RabbitMQ and declared queue: chat.logs.queue")
	return nil
}

func (r *RabbitMQ) PublishMessage(ctx context.Context, msg map[string]interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = r.channel.PublishWithContext(ctx,
		"",                // exchange
		"chat.logs.queue", // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	return err
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
