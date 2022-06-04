package blog

import (
	"strconv"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.canergulay/blogbackend/internal/server/models"
	"github.com/labstack/echo/v4"
)

func (h BlogManager) GetPostByIdHandler() endpoints.Endpoint {
	return endpoints.NewEndpoint("/post/:id", h.getPostByIdHandler, "GET")
}

func (h BlogManager) getPostByIdHandler(c echo.Context) error {
	id := c.Param("id")
	idParsed, err := strconv.Atoi(id)

	if err == nil {
		post := h.getPostsService.GetPostById(idParsed)
		if post != nil {
			response := models.Response{Status: 0, Data: post}
			c.JSON(200, response)
		} else {
			response := models.Response{Status: 1, Data: "There is not a blog post with that id."}
			c.JSON(404, response)
		}
	} else {
		c.JSON(500, err)
		// TODO: ADD LOGGING AND A PROPER ERROR RESPONSE WELL-STRUCTURED
	}

	return err
}
