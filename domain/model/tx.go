package model

import (
	"time"

	"gorm.io/gorm"
)

type Tx struct {
	ID            uint   `gorm:"primaryKey"`
	UserId        uint   `gorm:"not null"`
	AddressId     uint   `gorm:"not null"`
	TotalPrice    int    `gorm:"not null"`
	InvoiceCode   string `gorm:"type:varchar(255);not null"`
	PaymentMethod string `gorm:"type:varchar(50);not null"`
	Status        string `gorm:"type:varchar(50);not null"`
	CreatedAt     *time.Time
	UpdateAt      *time.Time
	DeletedAt     gorm.DeletedAt

	User     User       `gorm:"foreignKey:UserId"`
	Address  Address    `gorm:"foreignKey:AddressId"`
	DetailTx []DetailTx `gorm:"foreignKey:TxId"`
}

type DetailTx struct {
	ID         uint `gorm:"primaryKey"`
	TxId       uint `gorm:"not null"`
	StoreId    uint `gorm:"not null"`
	ProductId  uint `gorm:"not null"`
	Kuantity   int  `gorm:"not null"`
	TotalPrice int  `gorm:"not null"`
	CreatedAt  *time.Time
	UpdateAt   *time.Time
	DeletedAt  gorm.DeletedAt

	Tx      Tx      `gorm:"foreignKey:TxId"`
	Store   Store   `gorm:"foreignKey:StoreId"`
	Product Product `gorm:"foreignKey:ProductId"`
}
