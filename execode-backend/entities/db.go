package entities

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitPostgresQL(host string, user string, pw string, dbname string, port int) error {
	var err error
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%d sslmode=disable",
		host, user, pw, dbname, port)
	// dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", user, pw, host, port, dbname)

	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "execode.",
		},
	}
	db, err = gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		return err
	}

	return nil
}

func AutoMigrate() error {
	err := db.AutoMigrate(
		&User{},
		&Class{},
		&Lecture{},
		&Problem{},
		&LectureContent{},
		&ProblemContent{},
		&TestCase{},
		&PDFLecture{},
	)
	if db.Error != nil {
		return db.Error
	}
	return err
}

type EntityType interface {
	User | Class | Lecture | Problem
}

func GetAll[T EntityType, V any](model *T, data *[]V, limit, offset int) (int64, error) {
	var count int64
	err := db.Model(model).Count(&count).Scopes(Paginate(limit, offset)).Find(data).Error
	return count, err
}

func Create[T EntityType](data *T) error {
	err := db.Create(data).Error
	return err
}

func QueryOne[T EntityType, V any](structQuery *T, data *V) error {
	err := db.Model(structQuery).Where(structQuery).First(data).Error
	return err
}

func QueryMany[T EntityType, V any](structQuery *T, data *[]V, limit, offset int) (int64, error) {
	var count int64
	err := db.Model(structQuery).Where(structQuery).Count(&count).Scopes(Paginate(limit, offset)).Find(data).Error
	return count, err
}
