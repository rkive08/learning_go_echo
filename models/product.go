package models

import "time"

// type Product struct {
// 	ID         uint            `json:"id" gorm:"primaryKey"`
// 	Name       string          `json:"name"`
// 	Price      int             `json:"price"`
// 	CategoryID *uint           `json:"category_id"` // NULLABLE
// 	Category   CategoryProduct `json:"category" gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE;"`
// 	CreatedAt  time.Time
// }

type Product struct {
	ID         uint            `gorm:"primaryKey" json:"id"`
	Name       string          `json:"name"`
	Price      int             `json:"price"`
	CategoryID *uint           `json:"category_id"` // penting: nullable dan harus ada tag json
	Category   CategoryProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`
	CreatedAt  time.Time       `json:"created_at"`
}
