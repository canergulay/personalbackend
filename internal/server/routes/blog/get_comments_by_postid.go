package blog

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.canergulay/blogbackend/internal/server/endpoints"
	"github.com/labstack/echo/v4"
)

func (h BlogManager) GetCommentsByPostId() endpoints.Endpoint {
	return endpoints.NewEndpoint("/comment", h.getCommentsByPostId, "GET")
}

func (h BlogManager) getCommentsByPostId(c echo.Context) error {
	bodyMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&bodyMap)
	if err != nil {
		// TODO: ADD PROPER ERROR HANDLING.
		c.String(http.StatusInternalServerError, "something unexpected had happened, we are working on it")
		return err
	}
	postid, ok := bodyMap["post_id"]

	if !ok {
		errMessage := "you have to specify a post_id in the request body"
		c.String(http.StatusNotFound, errMessage)
		return errors.New(errMessage)
	}

	parsedPostId, isInt := postid.(int)

	if !isInt {
		errMessage := "what you have specified as post_id should be an integer... but you have specified" + reflect.TypeOf(postid).String()
		c.String(http.StatusBadRequest, errMessage)
		return errors.New(errMessage)
	}

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
