package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func getJwtToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := strings.Split(c.Get(fiber.HeaderAuthorization), "Bearer ")
	if len(tokenString) < 2 {
		return nil, fmt.Errorf("invalid token")
	}

	return parseJWT(tokenString[1])
}

func Protected(c *fiber.Ctx) error {
	token, err := getJwtToken(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return fiber.ErrUnauthorized
		}

		return c.Next()
	}

	return fiber.ErrUnauthorized
}

func ProtectedWeb(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Redirect("/admin/login")
	}

	token, err := parseJWT(cookie)
	if err != nil {
		return c.Redirect("/admin/login")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Redirect("/admin/login")
		}

		return c.Next()
	}

	return c.Redirect("/admin/login")
}
