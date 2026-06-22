package main

import (
	"log"
	"os"

	"github.com/b1g-nguyx/strangerchat-backend/internal/broker"
	"github.com/b1g-nguyx/strangerchat-backend/internal/chat"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
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

	// 2. Initialize Hub
	hub := chat.NewHub()
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

		client := &chat.Client{
			Hub:       hub,
			Conn:      c,
			Send:      make(chan []byte, 256),
			SessionID: sessionID,
			UserID:    userID,
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
