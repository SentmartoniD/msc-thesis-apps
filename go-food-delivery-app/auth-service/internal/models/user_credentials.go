package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	Admin    UserRole = "ADMIN"
	Customer UserRole = "CUSTOMER"
)

type UserCredentials struct {
	ID        string    `gorm:"type:varchar(36);primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	UserRole UserRole
	Email    string `gorm:"unique"`
	Password string
}

func (c *UserCredentials) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return nil
}

func (UserCredentials) TableName() string {
	return "user_credentials"
}
