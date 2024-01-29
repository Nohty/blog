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
	user := api.Group("/user")
	user.Put("/", middleware.ProtectedWeb, handler.UpdateUser)

	// API routes Blog
	blog := api.Group("/blog")
	blog.Get("/", middleware.ProtectedWeb, handler.GetBlogPosts)
	blog.Get("/:id", middleware.ProtectedWeb, handler.GetBlogPost)
	blog.Post("/", middleware.ProtectedWeb, handler.CreateBlogPost)
	blog.Put("/:id", middleware.ProtectedWeb, handler.UpdateBlogPost)
	blog.Delete("/:id", middleware.ProtectedWeb, handler.DeleteBlogPost)

	// Web routes Blog
	web.Get("/", handler.HomeHandler)
	web.Get("/blog", handler.BlogHandler)
	web.Get("/blog/:id", handler.BlogPostHandler)

	// // Web routes Admin
	admin := web.Group("/admin")
	admin.Get("/login", handler.AdminLoginHandler)
	admin.Get("/", middleware.ProtectedWeb, handler.AdminHandler)
}
