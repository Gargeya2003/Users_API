package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"

	"users_api/db/sqlc"
	"users_api/internal/handler"
	"users_api/internal/logger"
	"users_api/internal/middleware"
	"users_api/internal/repository"
	"users_api/internal/routes"
	"users_api/internal/service"
)

func main() {
	// ---------------- Logger ----------------
	zapLogger := logger.New()
	defer zapLogger.Sync()

	zapLogger.Info("starting users API")

	// ---------------- Configuration ----------------
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		zapLogger.Fatal("DB_DSN environment variable not set")
	}

	// ---------------- Database ----------------
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		zapLogger.Fatal("failed to open database", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		zapLogger.Fatal("failed to connect to database", zap.Error(err))
	}

	zapLogger.Info("database connected")

	// ---------------- Dependency Injection ----------------
	q := sqlc.New(db)
	repo := repository.NewUserRepository(q)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	// ---------------- Fiber App ----------------
	app := fiber.New()

	// Middleware (order matters)
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(zapLogger))

	// Routes
	routes.Register(app, h)

	// ---------------- Server ----------------
	port := ":8080"
	zapLogger.Info("server listening", zap.String("port", port))

	if err := app.Listen(port); err != nil {
		zapLogger.Fatal("server failed to start", zap.Error(err))
	}
}
