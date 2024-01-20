package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
