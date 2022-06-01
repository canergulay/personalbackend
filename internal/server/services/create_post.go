package services

import "github.canergulay/blogbackend/internal/server/routes/blog"

type CreatePost struct{}

type CreatePostHandler struct {
	blogManager *blog.BlogManager
}

func NewCreatePostHandler(blogManager *blog.BlogManager) CreatePostHandler {
	return CreatePostHandler{blogManager: blogManager}
}

func (cph CreatePostHandler) Create() int {
	createdPostID := cph.blogManager.CreateBlog()
	return createdPostID
}
