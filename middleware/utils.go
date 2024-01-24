package middleware

import (
	"fmt"
	"time"

	"github.com/Nohty/blog/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(config.JWT_SECRET))
}

func parseJWT(token string) (*jwt.Token, error) {
	tokenParsed, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return tokenParsed, nil
}
