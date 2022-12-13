package entities

import (
	"fmt"
	"os"
	"path"
)

type PDFLecture struct {
	ClassName string `gorm:"primaryKey"`
	FileName  string `gorm:"primaryKey"`
	Module    string
}

type Courses struct {
	ClassName string
}

const BasePDFLecturePath = "_local/pdf-lectures/"

func UploadPDFController(filename, classname, module, lectureFile string) error {
	filepath := path.Join(BasePDFLecturePath, filename)
	if err := os.WriteFile(filepath, []byte(lectureFile), 0644); err != nil {
		return err
	}
	if err := db.Create(PDFLecture{ClassName: classname, FileName: filename, Module: module}).Error; err != nil {
		return err
	}

	return nil
}

func GetCourses() ([]Courses, error) {
	courses := []Courses{}
	if err := db.Model(&PDFLecture{}).Select("DISTINCT ON (class_name) class_name").Find(&courses).Error; err != nil {
		return []Courses{}, err
	}
	fmt.Println(courses)
	return courses, nil
}
