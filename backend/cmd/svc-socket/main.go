package main

import (
	"log"
	"os"

	"github.com/b1g-nguyx/strangerchat-backend/internal/broker"
	chatws "github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/delivery/websocket"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/repository"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/usecase"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func main() {
	app := fiber.New()

	// 1. Initialize RabbitMQ
	rmqURL := os.Getenv("RABBITMQ_URL")
	if rmqURL == "" {
		rmqURL = "amqp://guest:guest@localhost:5672/"
	}
	if err := broker.InitRabbitMQ(rmqURL); err != nil {
		log.Printf("Warning: Failed to connect to RabbitMQ: %v", err)
	} else {
		defer broker.RMQ.Close()
	}

	// 1.5 Initialize Redis & Usecase
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{Addr: redisURL})
	redisRepo := repository.NewRedisRoomRepo(rdb)
	chatUsecase := usecase.NewChatUsecase(redisRepo)

	// 2. Initialize Hub
	hub := chatws.NewHub()
	go hub.Run()

	// 3. WebSocket Upgrade Middleware
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// 4. WebSocket Route
	app.Get("/ws/chat", websocket.New(func(c *websocket.Conn) {
		// In production, get SessionID and UserID from token (e.g. query param c.Query("token"))
		sessionID := c.Query("session_id", "default_session")
		userID := c.Query("user_id", "anonymous")

		client := &chatws.Client{
			Hub:       hub,
			Conn:      c,
			Send:      make(chan []byte, 256),
			SessionID: sessionID,
			UserID:    userID,
			Usecase:   chatUsecase,
		}

		client.Hub.Register <- client

		// Run writePump in a separate goroutine
		go client.WritePump()
		// Run readPump in the current goroutine (blocking)
		client.ReadPump()
	}))

	port := os.Getenv("SOCKET_PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Starting Socket Service on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
