package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uint      `gorm:"primaryKey:autoIncrement"`
	UUID         uuid.UUID `gorm:"type:uuid;not null"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Password     string    `gorm:"type:varchar(255);not null"`
	PhoneNumber  string    `gorm:"type:varchar(15);unique"`
	TanggalLahir *time.Time
	Gender       string `gorm:"type:varchar(20);not null"`
	Description  string `gorm:"type:text"`
	Work         string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(100);not null"`
	RoleID       uint   `gorm:"type:uint;not null"`
	CreatedAt    *time.Time
	UpdateAt     *time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Store   *Store    `gorm:"foreignKey:UserId"`
	Address []Address `gorm:"foreignKey:UserId"`
	Tx      []Tx      `gorm:"foreginKey:UserId"`
}
