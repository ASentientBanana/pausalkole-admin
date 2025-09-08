package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type EntityField struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid" json:"id"`
	Field     string         `json:"field"`
	Value     string         `json:"value"`
	IsVisible bool           `json:"is_visible"`
	EntityID  uuid.UUID      `json:"entity_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type EntityType string

const (
	RecipientEntity EntityType = "recipient"
	AgencyEntity    EntityType = "agency"
)

type Entity struct {
	CreatedAt time.Time      `json:"created_at"`
	Type      EntityType     `json:"type" binding:"required,oneof=recipient agency"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid" json:"id"`
	Fields    []EntityField  `gorm:"foreignKey:EntityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"fields"`
	Name      string         `json:"name"`
	UserID    string         `json:"user_id"`
}
