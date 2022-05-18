package repository

import "time"

type TodoCreateRepository struct {
	Uuid  string `gorm:"primaryKey;column:uuid"`
	Title string `gorm:"column:title"`
	// Status    enums.Status `"status" gorm:"column:status"`
	Image     string    `gorm:"column:image"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
	// UpdatedAt time.Time `gorm:"autoUpdateTime;column:created_at"`
}

type TodoGetRepository struct {
	Uuid  string `gorm:"primaryKey;column:uuid"`
	Title string `gorm:"column:title"`
	// Status    enums.Status `"status" gorm:"column:status"`
	Image     string    `gorm:"column:image"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
	// UpdatedAt time.Time `gorm:"autoUpdateTime;column:created_at"`
}
