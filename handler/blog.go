package handler

import (
	"github.com/Nohty/blog/database"
	"github.com/Nohty/blog/model"
	"github.com/Nohty/blog/utils"
	"github.com/gofiber/fiber/v2"
)

func GetBlogPosts(c *fiber.Ctx) error {
	var posts []model.BlogPost
	if err := database.DB.Find(&posts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Response(c, fiber.StatusOK, posts)
}

func GetBlogPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	var post model.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	return utils.Response(c, fiber.StatusOK, post)
}

func CreateBlogPost(c *fiber.Ctx) error {
	type NewBlogPost struct {
		Title     string `json:"title" validate:"required"`
		Content   string `json:"content" validate:"required"`
		Category  string `json:"category" validate:"required"`
		Image     string `json:"image" validate:"required"`
		Published bool   `json:"published"`
	}

	input := new(NewBlogPost)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	post := model.BlogPost{
		Title:     input.Title,
		Content:   input.Content,
		Category:  input.Category,
		Image:     input.Image,
		Published: input.Published,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Response(c, fiber.StatusCreated, post)
}

func UpdateBlogPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	type UpdateBlogPost struct {
		Title     string `json:"title" validate:"required"`
		Content   string `json:"content" validate:"required"`
		Category  string `json:"category" validate:"required"`
		Image     string `json:"image" validate:"required"`
		Published bool   `json:"published"`
	}

	input := new(UpdateBlogPost)

	if err := utils.ParseBodyAndValidate(c, input); err != nil {
		return err
	}

	var post model.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Category = input.Category
	post.Image = input.Image
	post.Published = input.Published

	if err := database.DB.Save(&post).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Response(c, fiber.StatusOK, post)
}

func DeleteBlogPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}

	var post model.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	if err := database.DB.Delete(&post).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Response(c, fiber.StatusNoContent, nil)
}
