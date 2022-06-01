package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"desc"`
	UserName    string `json:"username" gorm:"foreginKey:Username" `
	Active      bool   `json:"is_active"`
	Posts       []Post `json:"posts"`
}
