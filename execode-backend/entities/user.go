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
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	UserName  string          `json:"user_name"`
	FirstName string          `json:"first_name,omitempty"`
	LastName  string          `json:"last_name,omitempty"`
	Email     string          `json:"email"`
	Class     []APIClassBasic `json:"classes" gorm:"many2many:class_lecturer"`
}

func (u User) Basic() APIUserBasic {
	return APIUserBasic{u.ID, u.UserName, u.FirstName, u.LastName, u.Email}
}

func (u User) Advanced() APIUserAdvanced {
	return APIUserAdvanced{u.ID, u.CreatedAt, u.UpdatedAt, u.UserName, u.FirstName, u.LastName, u.Email, ClassList(u.Class).Basic()}
}

type UserList []User

func (us UserList) Advanced() []APIUserBasic {
	ret := make([]APIUserBasic, len(us))
	for i, u := range us {
		ret[i] = u.Basic()
	}
	return ret
}

func GetUsers(pageSize, page int) ([]APIUserBasic, int64, error) {
	users := []APIUserBasic{}
	count, err := GetAll(&User{}, &users, pageSize, page)
	return users, count, err
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

	err = Create(&user)
	if err != nil {
		return APIUserBasic{}, err
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
	structQuery := User{UserName: username}
	err := QueryOne(&structQuery, &user)
	if err != nil {
		return user, err
	}
	structQuery.ID = user.ID
	err = db.Model(&structQuery).Association("Class").Find(&user.Class)
	if err != nil {
		return user, err
	}
	return user, err
}

func QueryUsersByUsername(usernames []string, data *[]User) error {
	err := db.Table("execode.users").Where("user_name IN ?", usernames).Find(data).Error
	return err
}
