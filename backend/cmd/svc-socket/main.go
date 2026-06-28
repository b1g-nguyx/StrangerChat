package main

import (
	"context"
	"log"
	"os"

	"github.com/b1g-nguyx/strangerchat-backend/config"
	"github.com/b1g-nguyx/strangerchat-backend/internal/broker"
	"github.com/b1g-nguyx/strangerchat-backend/internal/common/jwt"
	chatws "github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/delivery/websocket"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/repository"
	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/usecase"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 0. Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	jwtManager := jwt.New(cfg.JWT.Secret, cfg.JWT.AccessTokenExpiry)

	app := fiber.New()

	// 1. Initialize RabbitMQ
	rmqURL := os.Getenv("RABBITMQ_URL")
	if rmqURL == "" && cfg.RMQ.URL != "" {
		rmqURL = cfg.RMQ.URL
	} else if rmqURL == "" {
		rmqURL = "amqp://guest:guest@localhost:5672/"
	}
	if err := broker.InitRabbitMQ(rmqURL); err != nil {
		log.Printf("Warning: Failed to connect to RabbitMQ: %v", err)
	} else {
		defer broker.RMQ.Close()
	}

	// 1.5 Initialize Redis & Usecase
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" && cfg.Redis.URL != "" {
		redisURL = cfg.Redis.URL
	} else if redisURL == "" {
		redisURL = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{Addr: redisURL})
	redisRepo := repository.NewRedisRoomRepo(rdb)
	chatUsecase := usecase.NewChatUsecase(redisRepo)

	// Start Matchmaking worker
	chatUsecase.StartMatchmakingWorker(context.Background())

	// 2. Initialize Hub
	hub := chatws.NewHub(redisRepo)
	go hub.Run(context.Background())

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
		// Get JWT token from query string
		token := c.Query("token")
		if token == "" {
			log.Println("WebSocket connection rejected: missing token")
			c.Close()
			return
		}

		// Parse token
		userID, err := jwtManager.ParseToken(token)
		if err != nil {
			log.Printf("WebSocket connection rejected: invalid token: %v", err)
			c.Close()
			return
		}

		client := &chatws.Client{
			Hub:       hub,
			Conn:      c,
			Send:      make(chan []byte, 256),
			SessionID: "",
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
