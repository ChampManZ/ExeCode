package models

import "gorm.io/gorm"

type Lecture struct {
	gorm.Model
	ClassID            int `gorm:"not null"`
	Class              Class
	LectureName        string `gorm:"size:60;not null;index"`
	LectureDescription string `gorm:"size:500"`
	LectureContent     LectureContent
}

type LectureContent struct {
	LectureID int    `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
}
