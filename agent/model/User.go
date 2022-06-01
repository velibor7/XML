package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `json:"username" grom:"unique"`
	Password string    `json:"password"`
	Role     Role      `json:"role"`
	Company  []Company `json:"companies" gorm:"-"`
}

type Role int

const (
	USER Role = iota
	OWNER
	ADMIN
)
