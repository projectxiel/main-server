package handlers

import (
	"projectxiel/data"

	"github.com/gofiber/fiber/v2"
)

// GetSinglePost godoc
// @Summary      Get a single post
// @Description  get post by slug
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        slug   path      string  true  "Post Slug"
// @Success      200  {object}  data.Post
// @Router       /post/{slug} [get]
func GetSinglePost(c *fiber.Ctx) error {
	slug := c.Params("slug")
	d := data.GetSinglePost(slug)
	return c.JSON(d)
}

// GetAllPosts godoc
// @Summary      Get all posts
// @Description  get posts
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        limit query      int false "Posts Limit"
// @Param 		 page query 		int false "Posts page"
// @Success      200  {object}  data.Post
// @Router       /posts [get]
func GetAllPosts(c *fiber.Ctx) error {
	m := c.Queries()
	limit := m["limit"]
	page := m["page"]
	d, err := data.GetAllPosts(limit, page)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(d)
}
