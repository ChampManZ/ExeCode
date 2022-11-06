package entities

import (
	"fmt"

	"github.com/ChampManZ/ExeCode/v2/entities/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

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
	DB, err = gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		return err
	}

	return nil
}

func AutoMigrate() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Class{},
		&models.Lecture{},
		&models.Problem{},
		&models.LectureContent{},
		&models.ProblemContent{},
		&models.TestCase{},
	)
	if DB.Error != nil {
		return DB.Error
	}
	return err
}
