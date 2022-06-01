package services

import (
	"github.canergulay/blogbackend/internal/server/routes/blog"
)

type SavePostService struct {
	blogManager *blog.BlogManager
}

func NewSavePostService(blogManager *blog.BlogManager) SavePostService {
	return SavePostService{blogManager: blogManager}
}

func (svph SavePostService) Save(post *blog.Post) *blog.Post {
	return svph.blogManager.SavePost(post)
}
