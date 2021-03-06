package models

import "time"

type Comment struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Text      string     `json:"text"`
	Owner     string     `json:"owner"`
	PostId    int        `json:"post_id" gorm:"index"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
