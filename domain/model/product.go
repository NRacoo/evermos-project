package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255);not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt

	Product []Product
}

type Product struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"type:varchar(255);not null"`
	Slug          string `gorm:"type:varchar(255);not null"`
	ResellerPrice int    `gorm:"not null"`
	Price         int    `gorm:"not null"`
	Stock         int    `gorm:"not null"`
	Description   string `gorm:"type:varchar(255);not null"`
	StoreId       uint   `gorm:"not null"`
	CategoryId    uint   `gorm:"not null"`
	CreatedAt     *time.Time
	UpdateAt      *time.Time
	DeletedAt     gorm.DeletedAt

	Store        Store          `gorm:"foreignKey:StoreId"`
	Category     Category       `gorm:"foreignKey:CategoryId"`
	ImageProduct []ImageProduct `gorm:"foreignKey:ProductId"`
}

type ImageProduct struct {
	ID        uint   `gorm:"primaryKey"`
	ProductId uint   `gorm:"not null"`
	Url       string `gorm:"type:varchar(255)"`
	CreatedAt *time.Time
	UpdateAt  *time.Time
	DeletedAt gorm.DeletedAt

	Product Product `gorm:"foreignKey:ProductId"`
}

type ProductLog struct {
	ID            uint   `gorm:"primaryKey"`
	ProductId     uint   `gorm:"not null"`
	StoreId       uint   `gorm:"not null"`
	NameProduct   string `gorm:"type:varchar(255);not null"`
	Slug          string `gorm:"type:varchar(255);not null"`
	ResellerPrice int    `gorm:"not null"`
	Price         int    `gorm:"not null"`
	Description   string `gorm:"type:text"`
	CreatedAt     *time.Time
	UpdateAt      *time.Time

	Product Product `gorm:"foreignKey:ProductId"`
	Store   Store   `gorm:"foreignKey:StoreID"`
}
