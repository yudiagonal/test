package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint64         `gorm:"primaryKey" json:"id"`
	NamaBarang  string         `gorm:"column:nama_barang" json:"nama_barang"`
	Harga       string         `gorm:"column:harga" json:"harga"`
	Jenis       string         `gorm:"column:jenis" json:"jenis"`
	MetaKeyword string         `gorm:"column:meta_keyword" json:"meta_keyword"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index,column:deleted_at" json:"deleted_at"`
}
