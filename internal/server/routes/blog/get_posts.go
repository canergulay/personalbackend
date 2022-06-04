package blog

import (
	"encoding/json"
	"io"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.canergulay/blogbackend/internal/server/models"
	"github.com/labstack/echo/v4"
)

func (h BlogManager) GetPostsEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/blog", h.getPostsHandler, "GET")
}

func (h BlogManager) getPostsHandler(c echo.Context) error {

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
