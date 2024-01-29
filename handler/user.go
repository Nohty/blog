package handler

import (
	"github.com/Nohty/blog/database"
	"github.com/Nohty/blog/model"
	"github.com/Nohty/blog/utils"
	"github.com/gofiber/fiber/v2"
)

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Name       string `json:"name" validate:"required"`
		Email      string `json:"email" validate:"required,email"`
		Profession string `json:"profession" validate:"required"`
		ImageUrl   string `json:"imageUrl" validate:"required"`
		Linkedin   string `json:"linkedin" validate:"required"`
		Github     string `json:"github" validate:"required"`
	}

	input := new(UpdateUserInput)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	var user model.User
	if err := database.DB.First(&user, database.USER_ID).Error; err != nil {
		return fiber.ErrNotFound
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Profession = input.Profession
	user.ImageUrl = input.ImageUrl
	user.Linkedin = input.Linkedin
	user.Github = input.Github

	if err := database.DB.Save(&user).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Response(c, fiber.StatusOK, user)
}
