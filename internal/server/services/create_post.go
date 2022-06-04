package services

import (
	"github.canergulay/blogbackend/internal/server/models"
	"gorm.io/gorm"
)

type CreatePost struct{}

type CreatePostService struct {
	DB *gorm.DB
}

func NewCreatePostService(db *gorm.DB) CreatePostService {
	return CreatePostService{DB: db}
}

func (cph CreatePostService) Create() int {
	post := models.Post{}
	cph.DB.Create(&post)
	return post.ID
}
