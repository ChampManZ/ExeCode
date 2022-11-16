package entities

import (
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ID               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index" swaggertype:"string" json:"-"`
	ClassName        string         `gorm:"size:50;not null;index;unique"`
	ClassDescription string         `gorm:"size:1000"`
	User             []User         `gorm:"many2many:class_lecturer" json:"-"`
	Lectures         []Lecture
} // @name ClassFields

type APIClassBasic struct {
	ID               uint   `json:"id"`
	ClassName        string `json:"class_name"`
	ClassDescription string `json:"class_description"`
} // @name ClassBasic

type APIClassAdvanced struct {
	ID               uint           `json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at" swaggertype:"string"`
	ClassName        string         `json:"class_name"`
	ClassDescription string         `json:"class_description"`
	User             []APIUserBasic `json:"users" gorm:"many2many:class_lecturer;joinForeignKey:class_id;References:id;joinReferences:user_id"`
} // @name ClassAdvanced

type ClassList []Class

func (c Class) Basic() APIClassBasic {
	return APIClassBasic{c.ID, c.ClassName, c.ClassDescription}
}

func (cs ClassList) Basic() []APIClassBasic {
	ret := make([]APIClassBasic, len(cs))
	for i, c := range cs {
		ret[i] = c.Basic()
	}
	return ret
}

func (c Class) Advanced() APIClassAdvanced {
	return APIClassAdvanced{
		ID:               c.ID,
		CreatedAt:        c.CreatedAt,
		UpdatedAt:        c.UpdatedAt,
		ClassName:        c.ClassName,
		ClassDescription: c.ClassDescription,
		User:             UserList(c.User).Advanced(),
	}
}

func GetClassByID(id uint, fields ...string) (Class, error) {
	class := Class{}
	class.ID = id
	var err error
	if err = db.Model(&class).Select(fields).Where(&class).First(&class).Error; err != nil {
		return Class{}, err
	}

	if err = db.Model(&class).Association("User").Find(&class.User); err != nil {
		return Class{}, err
	}
	return class, err
}

func DeleteClass(id uint) error {
	err := db.Delete(&Class{}, id).Error
	return err
}
