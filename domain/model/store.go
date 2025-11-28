package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	ID        uint   `gorm:"primaryKey"`
	UserId    uint   `gorm:"not null;unique"`
	Name      string `gorm:"type:varchar(255);not null"`
	ImageUrl  string `gorm:"type:varchar(255);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt

	User User `gorm:"foreignKey:UserId"`
}
