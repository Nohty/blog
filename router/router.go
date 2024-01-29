package router

import (
	"github.com/Nohty/blog/handler"
	"github.com/Nohty/blog/middleware"
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
	// user := api.Group("/user")
	// user.Post("/", handler.CreateUser)
	// user.Post("/refresh", handler.CreateBlogPostt)

	// API routes Blog
	blog := api.Group("/blog")
	blog.Get("/", middleware.Protected, handler.GetBlogPosts)
	blog.Get("/:id", middleware.Protected, handler.GetBlogPost)
	blog.Post("/", middleware.Protected, handler.CreateBlogPost)
	blog.Put("/:id", middleware.Protected, handler.UpdateBlogPost)
	blog.Delete("/:id", middleware.Protected, handler.DeleteBlogPost)

	// Web routes Blog
	web.Get("/", handler.HomeHandler)
	web.Get("/blog", handler.BlogHandler)
	web.Get("/blog/:id", handler.BlogPostHandler)

	// // Web routes Admin
	admin := web.Group("/admin")
	admin.Get("/login", handler.AdminLoginHandler)
	admin.Get("/", handler.AdminHandler)
}
