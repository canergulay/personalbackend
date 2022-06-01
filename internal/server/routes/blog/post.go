package blog

import "gorm.io/gorm"

type BlogManager struct {
	DB *gorm.DB
}

func NewBlogManager(db *gorm.DB) BlogManager {
	return BlogManager{DB: db}
}

func (bm BlogManager) CreateBlog() int {
	post := Post{}
	bm.DB.Create(&post)
	return post.ID
}
