package services

import (
	"github.canergulay/blogbackend/internal/server/models"
	"gorm.io/gorm"
)

type GetPostsService struct {
	DB *gorm.DB
}

// LAST ID WILL BE USED FOR PAGINATION PURPOSES
// WILL BE EVALUATED AS FIRST PAGE IF EMPTY
func NewGetPostService(db *gorm.DB) GetPostsService {
	return GetPostsService{DB: db}
}

func (gps *GetPostsService) GetPosts(lastID int) []models.Post {
	var posts []models.Post

	// IF THE LAST ID IS NOT SPECIFIED, IT WILL BE PASSED DOWN AS -1
	// WE WILL RETURN THE FIRST 10, WHICH REPRESENTS THE FIRST PAGE
	// IF IT IS SPECIFIED, IT MEANS WE ARE PAGINATING THROUGH A SPECIFIG PAGE.
	if lastID == -1 {
		gps.DB.Model(models.Post{}).Limit(10).Order("id desc").Find(&posts)
	} else {
		gps.DB.Model(models.Post{}).Where("id >", lastID).Limit(10).Order("id desc").Find(&posts)
	}
	return posts
}

func (gps *GetPostsService) GetPostById(id int) *models.Post {
	var post models.Post
	result := gps.DB.Model(models.Post{}).Where("id =", id).Find(&post)
	if result.Error != nil {
		return nil
	}
	return &post
}
