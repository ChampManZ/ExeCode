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
	)
	if db.Error != nil {
		return db.Error
	}
	return err
}
