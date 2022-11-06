package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserName       string         `gorm:"size:20;unique;index;not null"`
	FirstName      string         `gorm:"size:25;not null"`
	LastName       string         `gorm:"size:25;not null"`
	Email          string         `gorm:"size:80;unique;index;not null"`
	HashedPassword string         `gorm:"size:255"`
	Salt           string         `gorm:"size:25"`
	Lecture        []Lecture      `gorm:"many2many:class_lecturer"`
}
