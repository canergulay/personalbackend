package blog

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

func (h BlogManager) GetCommentsByPostIdEndpoint() endpoints.Endpoint {
	return endpoints.NewEndpoint("/comment", h.getCommentsByPostId, "POST")
}

func (h BlogManager) getCommentsByPostId(c echo.Context) error {
	bodyMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&bodyMap)
	fmt.Println("1")

	if err != nil {
		// TODO: ADD PROPER ERROR HANDLING.
		c.String(http.StatusInternalServerError, "something unexpected had happened, we are working on it")
		return err
	}
	fmt.Println("2")

	postid, ok := bodyMap["post_id"]
	if !ok {
		errMessage := "you have to specify a post_id in the request body"
		c.String(http.StatusNotFound, errMessage)
		return errors.New(errMessage)
	}

	fmt.Println("3")

	parsedPostId, err := strconv.Atoi(postid.(string))

	if err != nil {
		errMessage := "what you have specified as post_id should be an integer... but you have specified" + err.Error()
		c.String(http.StatusBadRequest, errMessage)
		fmt.Println(err)
		return errors.New(errMessage)
	}
	fmt.Println("4")

	comments, err := h.getCommentsService.GetCommentsByPostId(parsedPostId)
	if err != nil {
		// TODO: ADD PROPER ERROR HANDLING.
		fmt.Println(err)
		errMessage := "unexpected"
		c.String(http.StatusInternalServerError, errMessage)
		return errors.New(errMessage)
	}

	c.JSON(200, comments)
	return nil
}
