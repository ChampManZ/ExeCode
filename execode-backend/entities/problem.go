package entities

import (
	"time"

	"gorm.io/gorm"
)

type Problem struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index" swaggertype:"string"`
	ClassID        int            `gorm:"not null"`
	Class          Class
	ProblemName    string `gorm:"size:80;not null"`
	ProblemContent ProblemContent
}

type ProblemContent struct {
	ProblemID int        `gorm:"primaryKey"`
	Content   string     `gorm:"text;not null"`
	TestCases []TestCase `gorm:"foreignKey:ProblemID;references:ProblemID"`
}

type TestCase struct {
	ProblemID  int    `gorm:"primaryKey"`
	TestCaseID int    `gorm:"primaryKey"`
	Input      string `gorm:"type:text;not null"`
	Output     string `gorm:"type:text;not null"`
}
