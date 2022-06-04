package models

import "time"

type Post struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	NOV       int       `gorm:"default:0" json:"nov"`  // NUMBER OF VIEWS
	NOC       int       `gorm:"default:0" json:"noc" ` // NUMBER OF COMMENTS
}
