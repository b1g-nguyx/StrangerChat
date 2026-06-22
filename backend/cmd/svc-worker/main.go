package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ChatLogMessage struct {
	SessionID string `json:"session_id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
}

func main() {
	log.Println("Starting svc-worker...")

	// 1. Connect to Redis
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	// 2. Connect to RabbitMQ
	rmqURL := os.Getenv("RABBITMQ_URL")
	if rmqURL == "" {
		rmqURL = "amqp://guest:guest@localhost:5672/"
	}
	conn, err := amqp.Dial(rmqURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// 3. Declare Queue to ensure it exists
	q, err := ch.QueueDeclare(
		"chat.logs.queue", // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 4. Consume messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			var msg ChatLogMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			// Save to Redis as List: rpush chat_log:{session_id} {json_body}
			redisKey := "chat_log:" + msg.SessionID
			
			// Append message to the end of the list
			if err := rdb.RPush(ctx, redisKey, string(d.Body)).Err(); err != nil {
				log.Printf("Failed to save to Redis: %v", err)
				continue
			}

			// Set Time-To-Live (TTL) to 2 hours (7200 seconds) for new keys
			rdb.Expire(ctx, redisKey, 2*time.Hour)
			
			log.Printf("Saved message from %s (Session: %s) to Redis", msg.UserID, msg.SessionID)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
