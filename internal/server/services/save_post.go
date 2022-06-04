package services

import (
	"github.canergulay/blogbackend/internal/server/models"
	"gorm.io/gorm"
)

type SavePostService struct {
	DB *gorm.DB
}

func NewSavePostService(db *gorm.DB) SavePostService {
	return SavePostService{DB: db}
}

func (svph SavePostService) Save(post *models.Post) *models.Post {
	svph.DB.Save(post)
	return post
}
