package services

import (
	"fmt"

	"github.canergulay/blogbackend/internal/server/routes/blog"
)

type SavePostService struct {
	blogManager *blog.BlogManager
}

func NewSavePostService(blogManager *blog.BlogManager) SavePostService {
	return SavePostService{blogManager: blogManager}
}

func (svph SavePostService) Save(data interface{}) {
	savePost, ok := data.(blog.Post)
	if !ok {
		fmt.Println("data is not in the correct form !")
	}
	fmt.Println(savePost)
}
