package model

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model `json:"-"`
	Title      string `gorm:"not null" json:"title"`
	Content    string `gorm:"not null" json:"content"`
	Author     User   `gorm:"foreignKey:ID" json:"autor"`
	Category   string `gorm:"not null" json:"category"`
}
