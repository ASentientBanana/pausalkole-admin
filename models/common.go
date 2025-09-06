package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
