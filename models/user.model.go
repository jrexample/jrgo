package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64 `gorm:"primaryKey"`
	Username string
	Password string
	Name     string
}
