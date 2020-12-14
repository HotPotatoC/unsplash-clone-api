package main

import (
	"context"

	"github.com/HotPotatoC/unsplash-clone/pkg/database"
	"github.com/HotPotatoC/unsplash-clone/pkg/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_ = godotenv.Load()

	db := database.NewMongoHandler()

	server := server.NewFiberApp(ctx, db, ":5000", fiber.Config{
		Prefork: true,
	})

	server.App.Use(cors.New())
	server.App.Use(logger.New())
	server.App.Use(recover.New())

	server.SetupHandlers()
	server.Start()
}
