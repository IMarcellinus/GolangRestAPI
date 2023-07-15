package models

import (
	"time"
)

type Product struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(100)" json:"nama_product"`
	Deskripsi   string `gorm:"type:varchar(100)" json:"deskripsi"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
