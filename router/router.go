package router

import (
	"github.com/Nohty/blog/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Group routes
	api := app.Group("/api")
	web := app.Group("/")

	// API routes Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// API routes User
	user := api.Group("/user")
	user.Post("/", handler.CreateUser)

	// // API routes Blog
	// blog := api.Group("/blog")

	// Web routes Blog
	web.Get("/", handler.HomeHandler)

	// // Web routes Admin
	// admin := web.Group("/admin")
}
