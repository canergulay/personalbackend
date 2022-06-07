package services

import (
	"github.canergulay/blogbackend/internal/server/models"
	"gorm.io/gorm"
)

type GetCommentsService struct {
	DB *gorm.DB
}

func NewGetCommentsService(db *gorm.DB) GetCommentsService {
	return GetCommentsService{DB: db}
}

func (gcs *GetCommentsService) GetCommentsByPostId(postid int) (*[]models.Comment, error) {
	var comments []models.Comment
	result := gcs.DB.Model(models.Comment{}).Where("post_id = ?", postid).Find(&comments)
	if result.Error != nil {
		return &comments, result.Error
	}
	return &comments, nil
}
