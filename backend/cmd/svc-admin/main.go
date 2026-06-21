package main

import (
	"log"

	"github.com/b1g-nguyx/strangerchat-backend/config"
	"github.com/b1g-nguyx/strangerchat-backend/internal/app"
)

// @title StrangerChat Admin API
// @version 1.0
// @description API documentation for StrangerChat Admin Backend
// @host localhost:9090
// @BasePath /admin/v1
func main() {
	// 1. Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// 2. Run Admin Server
	app.RunAdmin(cfg)
}
