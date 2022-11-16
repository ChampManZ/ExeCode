package swaggercompat

import (
	"time"

	"gorm.io/gorm"
)

type ClassBasic struct {
	ID               uint   `json:"id"`
	ClassName        string `json:"class_name"`
	ClassDescription string `json:"class_description"`
} // @name ClassBasicNoRelation

type ClassAdvanced struct {
	ID               uint           `json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" swaggertype:"string"`
	ClassName        string         `json:"class_name"`
	ClassDescription string         `json:"class_description"`
} // @name ClassAdvancedNoRelation
