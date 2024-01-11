package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"projectxiel/data"
	"time"

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

// FetchYouTubeVideos godoc
// @Summary      Fetch YouTube videos
// @Description  Fetches a list of YouTube videos from a specific channel.
// @Tags         YouTube
// @Accept       json
// @Produce      json
// @Success      200  {object}  data.YouTubeResponse  "A list of YouTube videos"
// @Router       /youtube/videos [get]
func FetchYouTubeVideos(c *fiber.Ctx) error {
	cacheFile := "cache/youtube-videos.json"
	cacheURL := "/static/youtube-videos.json" // URL path for the static file

	// Check if cached file exists and is valid
	if fileInfo, err := os.Stat(cacheFile); err == nil {
		if time.Since(fileInfo.ModTime()) < time.Hour*24 {
			// Redirect to static route
			return c.Redirect(cacheURL)
		}
	}
	var ChannelID = os.Getenv("CHANNEL_ID")
	var YtApiKey = os.Getenv("YT_API_KEY")

	// Set up the YouTube API request
	youtubeAPIURL := fmt.Sprintf("https://youtube.googleapis.com/youtube/v3/search?part=snippet&channelId=%s&maxResults=50&key=%s", ChannelID, YtApiKey)

	// Make the request to the YouTube API
	res, err := http.Get(youtubeAPIURL)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch data from YouTube",
		})
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read YouTube response",
		})
	}

	// Parse the response body as JSON
	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		// Handle JSON parsing error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse YouTube response as JSON",
		})
	}
	// Write response to cache file
	os.WriteFile(cacheFile, body, 0644)
	// Forward the parsed JSON response
	return c.JSON(data)
}
