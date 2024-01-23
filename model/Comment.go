package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model `json:"-"`
	Content    string `gorm:"not null" json:"content"`
	Name       string `gorm:"not null" json:"name"`
	Email      string `gorm:"not null" json:"email"`
	BlogPost   uint   `gorm:"not null" json:"-"`
}
