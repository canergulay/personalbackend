package blog

import (
	"github.canergulay/blogbackend/internal/server/services"
)

type BlogManager struct {
	savePostService    *services.SavePostService
	getPostsService    *services.GetPostsService
	getCommentsService *services.GetCommentsService
}

func NewBlogManager(sps *services.SavePostService, gps *services.GetPostsService, gcs *services.GetCommentsService) *BlogManager {
	return &BlogManager{savePostService: sps, getPostsService: gps, getCommentsService: gcs}
}
