package models

import (
	"gorm.io/gorm"
	"time"
)

type Model interface {
	GetID() int64
	BeforeCreate(tx *gorm.DB) error
	BeforeUpdate(tx *gorm.DB) error
}

type BaseModel struct {
	Id        int64     `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *BaseModel) GetID() int64 {
	return m.Id
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	now = now.Round(time.Second)
	if m.CreatedAt.IsZero() {
		tx.Statement.SetColumn("created_at", now)
	}
	if m.UpdatedAt.IsZero() {
		tx.Statement.SetColumn("updated_at", now)
	}
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now()
	now = now.Round(time.Second)
	tx.Statement.SetColumn("UpdatedAt", now)
	return nil
}
