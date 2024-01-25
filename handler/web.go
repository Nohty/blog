package handler

import (
	"github.com/Nohty/blog/database"
	"github.com/Nohty/blog/model"
	"github.com/Nohty/blog/utils"
	"github.com/Nohty/blog/view/pages"
	"github.com/gofiber/fiber/v2"
)

func HomeHandler(c *fiber.Ctx) error {
	var user model.User
	if err := database.DB.First(&user, database.USER_ID).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Render(c, pages.Home(user))
}
