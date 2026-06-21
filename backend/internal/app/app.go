package app

import (
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"

	"github.com/b1g-nguyx/strangerchat-backend/config"
	_ "github.com/b1g-nguyx/strangerchat-backend/docs"
	clientv1auth "github.com/b1g-nguyx/strangerchat-backend/internal/controller/restapi/client/v1/auth"
	"github.com/b1g-nguyx/strangerchat-backend/internal/repo/persistent"
	"github.com/b1g-nguyx/strangerchat-backend/internal/usecase/user"
	"github.com/b1g-nguyx/strangerchat-backend/pkg/jwt"
)

// RunChat initializes and starts the Client/Chat application.
func RunChat(cfg *config.Config) {
	// 1. Initialize Database using secure configuration
	db, err := sql.Open("postgres", cfg.PG.URL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// 3. Initialize shared libraries with environment variables
	jwtManager := jwt.New(cfg.JWT.Secret, cfg.JWT.TokenExpiry)

	// 4. Initialize Repository and Usecase (Dependency Injection)
	userRepo := persistent.NewUserRepo(db)
	userUseCase := user.New(userRepo, jwtManager)

	// 5. Initialize Web Server
	fiberApp := fiber.New()
	fiberApp.Use(cors.New())

	// 6. Configure Routes
	if cfg.Swagger.Enabled {
		fiberApp.Get("/swagger/*", swagger.HandlerDefault)
	}
	
	v1Group := fiberApp.Group("/v1")
	clientv1auth.NewAuthRoutes(v1Group, userUseCase)

	// 7. Run Server in a separate goroutine
	go func() {
		log.Printf("Client Server is running on port %s", cfg.HTTP.Port)
		if err := fiberApp.Listen(":" + cfg.HTTP.Port); err != nil {
			log.Fatalf("Fiber listen error: %v", err)
		}
	}()

	// 8. Graceful Shutdown (Wait for OS signals like SIGINT, SIGTERM)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server gracefully...")
	if err := fiberApp.Shutdown(); err != nil {
		log.Fatalf("Fiber shutdown error: %v", err)
	}
	log.Println("Server exited properly")
}

// RunAdmin initializes and starts the Admin application.
func RunAdmin(cfg *config.Config) {
	// 1. Initialize Database using secure configuration
	db, err := sql.Open("postgres", cfg.PG.URL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// 2. Initialize shared libraries
	jwtManager := jwt.New(cfg.JWT.Secret, cfg.JWT.TokenExpiry)

	// 3. Initialize Repository and Usecase
	userRepo := persistent.NewUserRepo(db)
	userUseCase := user.New(userRepo, jwtManager)

	// 4. Initialize Web Server
	fiberApp := fiber.New()
	fiberApp.Use(cors.New())

	// 5. Configure Admin Routes (Placeholder for now)
	// TODO: Add admin-specific routes here using adminGroup
	// adminGroup := fiberApp.Group("/admin/v1")
	_ = userUseCase // prevent unused variable error until admin routes use it

	// 6. Run Server (Using a different port for Admin to avoid collision with Client if run on same machine)
	// In production, this should ideally be driven by a separate environment variable (e.g. ADMIN_HTTP_PORT)
	adminPort := cfg.HTTP.Port

	go func() {
		log.Printf("Admin Server is running on port %s", adminPort)
		if err := fiberApp.Listen(":" + adminPort); err != nil {
			log.Fatalf("Fiber listen error: %v", err)
		}
	}()

	// 7. Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down admin server gracefully...")
	if err := fiberApp.Shutdown(); err != nil {
		log.Fatalf("Fiber shutdown error: %v", err)
	}
	log.Println("Admin Server exited properly")
}
