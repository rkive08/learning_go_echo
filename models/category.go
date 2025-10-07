package models

import "time"

type CategoryProduct struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	// Products  []Product `gorm:"foreignKey:CategoryID"`
}

// jika nama tabel tidak sesuai dengan nama struct (pluralize)
// func (CategoryProduct) TableName() string {
// 	return "categories"
// }
