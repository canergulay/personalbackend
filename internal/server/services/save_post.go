package services

import (
	"fmt"

	"github.canergulay/blogbackend/internal/server/routes/blog"
)

type SavePostHandler struct {
	blogManager *blog.BlogManager
}

func NewSavePostHandler(blogManager *blog.BlogManager) SavePostHandler {
	return SavePostHandler{blogManager: blogManager}
}

func (svph SavePostHandler) Save(data interface{}) {
	savePost, ok := data.(blog.Post)
	if !ok {
		fmt.Println("data is not in the correct form !")
	}
	fmt.Println(savePost)
}
