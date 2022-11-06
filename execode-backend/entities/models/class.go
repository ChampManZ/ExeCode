package models

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	ClassName        string `gorm:"size:20;not null;index"`
	ClassDescription string `gorm:"size:1000"`
	User             []User `gorm:"many2many:class_lecturer"`
}
