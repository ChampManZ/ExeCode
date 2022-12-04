package entities

import (
	"time"

	"gorm.io/gorm"
)

type Lecture struct {
	ID                 uint `gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index" swaggertype:"string"`
	ClassID            int            `gorm:"not null;uniqueIndex:idx_class_lecture_name"`
	Class              Class
	LectureName        string `gorm:"size:60;not null;uniqueIndex:idx_class_lecture_name"`
	LectureDescription string `gorm:"size:500"`
	LectureContent     LectureContent
}

type LectureContent struct {
	LectureID uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
}

type APILectureBasic struct {
	ID                 uint   `json:"id"`
	ClassID            int    `json:"class_id"`
	LectureName        string `json:"lecture_name"`
	LectureDescription string `json:"lecture_description"`
}

type APILectureAdvanced struct {
	ID                 uint `json:"id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	ClassID            int    `json:"class_id"`
	LectureName        string `json:"lecture_name"`
	LectureDescription string `json:"lecture_description"`
	Content            string `json:"lecture_content"`
}

func (l Lecture) Basic() APILectureBasic {
	return APILectureBasic{l.ID, l.ClassID, l.LectureName, l.LectureDescription}
}

func (l Lecture) Advanced() APILectureAdvanced {
	return APILectureAdvanced{l.ID, l.CreatedAt, l.UpdatedAt, l.ClassID, l.LectureName, l.LectureDescription, l.LectureContent.Content}
}

type LectureList []Lecture

func (ls LectureList) Basic() []APILectureBasic {
	ret := make([]APILectureBasic, len(ls))
	for i, u := range ls {
		ret[i] = u.Basic()
	}
	return ret
}

func GetLectureByID(lectureID uint) (Lecture, error) {
	query := Lecture{ID: lectureID}
	err := db.Preload("LectureContent").Where(&query).First(&query).Error
	return query, err
}

func DeleteLecture(id uint) error {
	if err := db.Unscoped().Delete(&Lecture{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func (l *Lecture) BeforeDelete(tx *gorm.DB) (err error) {
	err = db.Delete(&LectureContent{LectureID: l.ID}).Error
	return
}

func GetClassLecturesByID(cid uint) (lectures []Lecture, err error) {
	class := Class{ID: cid}
	err = db.Model(&class).Association("Lectures").Find(&lectures)
	return
}
