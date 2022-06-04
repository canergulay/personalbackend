package blog

import (
	"encoding/json"
	"io"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.canergulay/blogbackend/internal/server/models"
	"github.canergulay/blogbackend/internal/server/services"
	"github.com/labstack/echo/v4"
)

type BlogManager struct {
	savePostService *services.SavePostService
	getPostsService *services.GetPostsService
}

func NewBlogManager(sps *services.SavePostService, gps *services.GetPostsService) *BlogManager {
	return &BlogManager{savePostService: sps, getPostsService: gps}
}

func (h BlogManager) GetBlogsHandler(c echo.Context) error {

	body := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&body)
	if err != io.EOF {
		c.String(500, "an unexpected error has occured"+err.Error()+"\n")
	}

	lastID := body["lastID"]

	lastIDParsed, ok := lastID.(int)

	var posts []models.Post

	if ok {
		posts = h.getPostsService.GetPosts(lastIDParsed)

	} else {
		posts = h.getPostsService.GetPosts(-1)

	}

	c.JSON(200, posts)

	return err
}

func (h BlogManager) GetBlogsEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/blog", h.GetBlogsHandler, "GET")
}
