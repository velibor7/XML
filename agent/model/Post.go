package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Description string `json:"desc"`
	CompanyId   uint   `json:"company_id"`
}
