package handler

import (
	"slices"
	"strings"

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

func BlogHandler(c *fiber.Ctx) error {
	category := c.Query("category")
	search := c.Query("search")

	var user model.User
	if err := database.DB.First(&user, database.USER_ID).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	var posts []model.BlogPost
	if err := database.DB.Where("published = ?", true).Find(&posts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	var filteredPosts []model.BlogPost

	if category != "" {
		for _, post := range posts {
			if post.Category == category {
				filteredPosts = append(filteredPosts, post)
			}
		}
	} else {
		filteredPosts = posts
	}

	if search != "" {
		var filteredSPosts []model.BlogPost
		for _, post := range filteredPosts {
			if strings.Contains(post.Title, search) {
				filteredSPosts = append(filteredPosts, post)
			}
		}

		filteredPosts = filteredSPosts
	}

	if category == "" && search == "" {
		filteredPosts = posts
	}

	var categories []string
	for _, post := range posts {
		if !slices.Contains(categories, post.Category) {
			categories = append(categories, post.Category)
		}
	}

	return utils.Render(c, pages.Blog(user, filteredPosts, search, category, categories))
}

func BlogPostHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Redirect("/blog")
	}

	var user model.User
	if err := database.DB.First(&user, database.USER_ID).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	var post model.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return fiber.ErrNotFound
	}

	return utils.Render(c, pages.Post(user, post))
}

func AdminHandler(c *fiber.Ctx) error {
	selected := c.Query("select")

	if selected == "profile" {
		var user model.User
		if err := database.DB.First(&user, database.USER_ID).Error; err != nil {
			return fiber.ErrInternalServerError
		}

		return utils.Render(c, pages.Admin(user, nil, true))
	}

	var posts []model.BlogPost
	if err := database.DB.Find(&posts).Error; err != nil {
		return fiber.ErrInternalServerError
	}

	return utils.Render(c, pages.Admin(model.User{}, posts, false))
}

func AdminLoginHandler(c *fiber.Ctx) error {
	return utils.Render(c, pages.Login())
}

func AdminNewHandler(c *fiber.Ctx) error {
	return utils.Render(c, pages.CreatePost(model.BlogPost{}))
}

func AdminEditHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Redirect("/admin")
	}

	var post model.BlogPost
	if err := database.DB.First(&post, id).Error; err != nil {
		return c.Redirect("/admin")
	}

	return utils.Render(c, pages.CreatePost(post))
}
