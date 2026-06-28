package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"database/sql"

	"github.com/b1g-nguyx/strangerchat-backend/internal/features/chat/repository"
	"github.com/b1g-nguyx/strangerchat-backend/internal/core/entity"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo/persistent"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ReportEventMessage struct {
	ReporterID string   `json:"reporter_id"`
	ReportedID string   `json:"reported_id"`
	RoomID     string   `json:"room_id"`
	Reason     string   `json:"reason"`
	ChatLogs   []string `json:"chat_logs"`
}

func main() {
	log.Println("Starting svc-worker...")

	// 1. Connect to PostgreSQL
	pgURL := os.Getenv("POSTGRES_URL")
	if pgURL == "" {
		pgURL = "postgres://user:myAwEsOm3pa55%40w0rd@localhost:5432/db?sslmode=disable"
	}
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	log.Println("Connected to Postgres successfully")

	baseRepo := persistent.NewBaseRepo(db)
	reportRepo := repository.NewPostgresReportRepo(baseRepo)

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
		"report.events.queue", // name
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
			var msg ReportEventMessage
			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			ctx := context.Background()
			reportID := uuid.New().String()
			
			// 1. Save Report
			report := entity.Report{
				BaseEntity: entity.BaseEntity{
					ID:        reportID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				ReporterID: msg.ReporterID,
				ReportedID: msg.ReportedID,
				Reason:     msg.Reason,
				Status:     "PENDING",
			}
			if err := reportRepo.CreateReport(ctx, report); err != nil {
				log.Printf("Failed to create report: %v", err)
				continue
			}

			// 2. Save Evidence
			chatLogsJSON, _ := json.Marshal(msg.ChatLogs)
			evidence := entity.ReportEvidence{
				BaseEntity: entity.BaseEntity{
					ID:        uuid.New().String(),
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				ReportID: reportID,
				RoomID:   msg.RoomID,
				ChatLogs: chatLogsJSON,
			}
			if err := reportRepo.SaveEvidence(ctx, evidence); err != nil {
				log.Printf("Failed to save evidence: %v", err)
				continue
			}

			log.Printf("Saved report and evidence for Reporter: %s, Reported: %s", msg.ReporterID, msg.ReportedID)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
