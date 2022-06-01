package services

import "github.canergulay/blogbackend/internal/server/routes/blog"

type CreatePost struct{}

type CreatePostService struct {
	blogManager *blog.BlogManager
}

func NewCreatePostService(blogManager *blog.BlogManager) CreatePostService {
	return CreatePostService{blogManager: blogManager}
}

func (cph CreatePostService) Create() int {
	createdPostID := cph.blogManager.CreateBlog()
	return createdPostID
}
