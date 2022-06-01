package handlers

import (
	"github.canergulay/blogbackend/internal/server/endpoints/routes/blog"
	"github.com/gorilla/websocket"
)

type CreatePostRespone struct {
	ID int `json:"id"`
}

type CreatePost struct{}

type CreatePostHandler struct {
	blogManager *blog.BlogManager
}

func NewCreatePostHandler(blogManager *blog.BlogManager) CreatePostHandler {
	return CreatePostHandler{blogManager: blogManager}
}

func (cph CreatePostHandler) Create(conn *websocket.Conn) {
	createdPostID := cph.blogManager.CreateBlog()
	response := CreatePostRespone{ID: createdPostID}
	conn.WriteJSON(response)
}
