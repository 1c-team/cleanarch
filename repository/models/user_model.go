package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int64
	Username  string
	Password  string
	Email     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
