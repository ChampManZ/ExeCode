package entities

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index" swaggertype:"string"`
	UserName       string         `gorm:"size:20;unique;index;not null"`
	FirstName      string         `gorm:"size:25;not null"`
	LastName       string         `gorm:"size:25;not null"`
	Email          string         `gorm:"size:80;unique;index;not null"`
	HashedPassword string         `gorm:"size:255" json:"-"`
	Salt           string         `gorm:"size:25" json:"-"`
	Class          []Class        `gorm:"many2many:class_lecturer"`
} // @name UserFields

type APIUserBasic struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
} // @name UserBasic

type APIUserAdvanced struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserName  string    `json:"user_name"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Email     string    `json:"email"`
	Class     []Class   `json:"classes" gorm:"many2many:class_lecturer"`
} // @name UserAdvanced

func GetUsers(pageSize, page int) ([]APIUserBasic, int64, error) {
	users := []APIUserBasic{}
	var count int64
	err := db.Model(&User{}).Scopes(Paginate(pageSize, page)).Find(&users).Count(&count).Error
	totalPages := count/int64(pageSize) + int64(page)
	return users, totalPages, err
}

func CreateUser(username, firstname, lastname, email, password string) (APIUserBasic, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return APIUserBasic{}, err
	}
	user := User{
		UserName:       username,
		FirstName:      firstname,
		LastName:       lastname,
		Email:          email,
		HashedPassword: string(hash),
	}

	result := db.Create(&user)
	if result.Error != nil {
		return APIUserBasic{}, result.Error
	}
	return APIUserBasic{
		user.ID,
		user.UserName,
		user.FirstName,
		user.LastName,
		user.Email,
	}, nil
}

func GetUserByUsername(username string) (APIUserAdvanced, error) {
	user := APIUserAdvanced{}
	structQuery := &User{UserName: username}
	err := db.Model(structQuery).Where(structQuery).First(&user).Error
	return user, err
}
