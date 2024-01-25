package main

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Nohty/blog/config"
	"github.com/Nohty/blog/database"
	"github.com/Nohty/blog/middleware"
	"github.com/Nohty/blog/router"
	"github.com/Nohty/blog/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed assets
var embedDirAssets embed.FS

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	database.Connect()

	router.SetupRoutes(app)

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirAssets),
		PathPrefix: "assets",
		Browse:     true,
	}))

	app.Use(logger.New())
	app.Use(middleware.NotFoundMiddleware)

	if app.Listen(fmt.Sprintf(":%s", config.PORT)) != nil {
		slog.Error("Error when starting server")
		os.Exit(1)
	}
}
