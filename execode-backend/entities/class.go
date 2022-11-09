package entities

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index" swaggertype:"string"`
	ClassName        string         `gorm:"size:20;not null;index"`
	ClassDescription string         `gorm:"size:1000"`
	User             []User         `gorm:"many2many:class_lecturer"`
} // @name ClassFields
