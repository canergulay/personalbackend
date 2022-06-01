package blog

import "time"

type Post struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Content   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
