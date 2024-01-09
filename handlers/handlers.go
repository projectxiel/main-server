package handlers

import (
	"projectxiel/data"

	"github.com/gofiber/fiber/v2"
)

func GetSinglePost(c *fiber.Ctx) error {
	slug := c.Params("slug")
	d := data.GetSinglePost(slug)
	return c.JSON(d)
}
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
