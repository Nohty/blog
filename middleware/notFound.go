package middleware

import (
	"net/http"

	"github.com/Nohty/blog/utils"
	"github.com/Nohty/blog/view/pages"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func NotFoundMiddleware(c *fiber.Ctx) error {
	return utils.Render(c, pages.NotFound(), templ.WithStatus(http.StatusNotFound))
}
