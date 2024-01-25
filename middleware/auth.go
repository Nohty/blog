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

		c.Next()
	}

	return fiber.ErrUnauthorized
}
