package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Nohty/blog/config"
	"github.com/Nohty/blog/database"
	"github.com/Nohty/blog/router"
	"github.com/Nohty/blog/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	database.Connect()

	app.Use(logger.New())
	router.SetupRoutes(app)

	if app.Listen(fmt.Sprintf(":%s", config.PORT)) != nil {
		slog.Error("Error when starting server")
		os.Exit(1)
	}
}
