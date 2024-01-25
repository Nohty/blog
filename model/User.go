package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	Name       string `gorm:"not null" json:"name"`
	Email      string `gorm:"uniqueIndex;not null" json:"email"`
	Password   string `gorm:"not null" json:"-"`
	Profession string `gorm:"not null" json:"profession"`
	ImageUrl   string `gorm:"not null" json:"imag_url"`
	Linkedin   string `gorm:"not null" json:"linkedin"`
	Github     string `gorm:"not null" json:"github"`
}
