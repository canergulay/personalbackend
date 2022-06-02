package services

import "github.canergulay/blogbackend/internal/server/routes/blog"

type GetPostsService struct {
	blogManager *blog.BlogManager
}

// LAST ID WILL BE USED FOR PAGINATION PURPOSES
// WILL BE EVALUATED AS FIRST PAGE IF EMPTY
func NewGetPostService(bm *blog.BlogManager) GetPostsService {
	return GetPostsService{blogManager: bm}
}

func (gps *GetPostsService) GetPosts(lastID int) []blog.Post {
	posts := gps.blogManager.GetPosts(lastID)
	return posts
}
