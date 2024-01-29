package handler

import (
	"time"

	"github.com/Nohty/blog/database"
	"github.com/Nohty/blog/middleware"
	"github.com/Nohty/blog/model"
	"github.com/Nohty/blog/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,ascii"`
	}

	input := new(LoginInput)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	var user model.User
	if err := database.DB.First(&user, database.USER_ID).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return fiber.ErrUnauthorized
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "jwt"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 24)
	c.Cookie(cookie)

	return utils.Response(c, fiber.StatusOK, token)
}
