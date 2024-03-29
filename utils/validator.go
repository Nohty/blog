package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Validate(payload interface{}) *fiber.Error {
	if err := validate.Struct(payload); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()))
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}

	return nil
}

func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}
