package model

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID          uint   `gorm:"primaryKey"`
	UserId      uint   `gorm:"not null"`
	Title       string `gorm:"type:varchar(255);not null"`
	Name        string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(15);not null"`
	Detail      string `gorm:"type:varchar(255);not null"`
	CreatedAt   *time.Time
	UpdateAt    *time.Time
	DeletedAt   gorm.DeletedAt

	User User `gorm:"foreignKey:UserId"`
}
