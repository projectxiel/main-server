package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

// GetCurrentProjects godoc
// @Summary      Get all current projects
// @Description  get current projects
// @Tags         Current Project
// @Accept       json
// @Produce      json
// @Param        limit query      int false "Projects Limit"
// @Param 		 page query 		int false "Projects page"
// @Success      200  {object}  data.CurrentProject
// @Router       /current-projects [get]
func GetCurrentProjects(c *fiber.Ctx) error {
	m := c.Queries()
	limit := m["limit"]
	page := m["page"]
	d, err := data.GetCurrentProjects(limit, page)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(d)
}

// SearchPosts godoc
// @Summary      Search Posts by title
// @Description  get posts containing title
// @Tags         Post
// @Accept       json
// @Produce      json
// @Param        title query     string true "Post title"
// @Param        limit query      int false "Posts Limit"
// @Param 		 page query 		int false "Posts page"
// @Success      200  {object}  data.Post
// @Router       /posts/search [get]
func SearchPosts(c *fiber.Ctx) error {
	m := c.Queries()
	title := m["title"]
	limit := m["limit"]
	page := m["page"]
	d, err := data.SearchPosts(title, limit, page)
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

	// Check if cached file exists and is valid
	if fileInfo, err := os.Stat(cacheFile); err == nil {
		if time.Since(fileInfo.ModTime()) < time.Hour*24 {
			// Redirect to static route
			return c.SendFile(cacheFile)
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

func init() {
	data, err := data.GetFileNames()
	if err != nil {
		log.Fatal(err)
	}
	filenameCache = sliceToKeyMap(data)
}

var filenameCache map[string]struct{}

// GetBlobData godoc
// @Summary      Serves static files
// @Description  Retrieves a static binary file from Blob Storage or cache, and serves it
// @Tags         Static
// @Accept       */*
// @Produce      */*
// @Param        blobname   path      string  true  "Filename"
// @Success      200  {file}  file "Binary File Content"
// @Router       /{blobname} [get]
func GetBlobData(c *fiber.Ctx) error {
	blobname := c.Params("blobname")
	cacheFile := "cache/" + blobname

	// Check if cached file exists and is valid
	if fileInfo, err := os.Stat(cacheFile); err == nil {
		if time.Since(fileInfo.ModTime()) < time.Hour*24 {
			// Redirect to static route
			return c.SendFile(cacheFile)
		}
	}
	if _, exists := filenameCache[blobname]; !exists {
		return c.Status(404).JSON(fiber.Map{
			"error": "Blob doesn't exist ",
		})
	}
	var err error
	if err = data.GetBlobData(blobname); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve Blob: " + blobname,
		})
	}
	return c.SendFile(cacheFile)
}

func sliceToKeyMap(slice []*data.Filename) map[string]struct{} {
	result := make(map[string]struct{})
	for _, item := range slice {
		if item != nil { // Ensure the pointer is not nil
			result[item.Name] = struct{}{}
		}
	}
	return result
}

var lastupdated time.Time

func UpdateCache(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	expected := "Bearer " + os.Getenv("CACHE_UPDATE_TOKEN")
	if expected == "" {
		log.Fatal("CACHE_UPDATE_TOKEN is not set")
	}
	if token != expected {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	filenames, err := data.GetFileNames()
	if err != nil {
		return err
	}

	filenameCache = sliceToKeyMap(filenames)

	for _, file := range filenames {
		if file.LastModified.After(lastupdated) {
			if err = data.GetBlobData(file.Name); err != nil {
				log.Println("Failed to retrieve blob:", file.Name)
				continue
			}
			log.Println(file.Name, "fetched")
		}
	}

	log.Println("Cache updated successfully")
	lastupdated = time.Now()
	return c.JSON(filenames)
}
