package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;type:uuid"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	//Files    []Files   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
