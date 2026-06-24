package main

import (
	"log"

	"github.com/b1g-nguyx/strangerchat-backend/config"
	"github.com/b1g-nguyx/strangerchat-backend/internal/app"
)

// @title StrangerChat Client API
// @version 1.0
// @description API documentation for StrangerChat Client Backend
// @host localhost:8080
// @BasePath /v1
func main() {
	// 1. Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// 2. Run
	app.RunClient(cfg)
}



