package model

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model `json:"-"`
	Title      string `gorm:"not null" json:"title"`
	Content    string `gorm:"not null" json:"content"`
	Category   string `gorm:"not null" json:"category"`
	Image      string `gorm:"not null" json:"image"`
}
