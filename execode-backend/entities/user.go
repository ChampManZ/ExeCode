package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" swaggertype:"string"`
	UserName  string         `gorm:"size:20;unique;index;not null"`
	Name      string         `gorm:"size:25;not null"`
	Email     string         `gorm:"size:80;unique;index;not null"`
	// HashedPassword string         `gorm:"size:255" json:"-"`
	// Salt           string         `gorm:"size:25" json:"-"`
	Class []Class `gorm:"many2many:class_lecturer"`
} // @name UserFields

type APIUserBasic struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
} // @name UserBasic

type APIUserAdvanced struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	UserName  string          `json:"user_name"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Class     []APIClassBasic `json:"classes" gorm:"many2many:class_lecturer"`
}

func (u User) Basic() APIUserBasic {
	return APIUserBasic{u.ID, u.UserName, u.Name, u.Email}
}

func (u User) Advanced() APIUserAdvanced {
	return APIUserAdvanced{u.ID, u.CreatedAt, u.UpdatedAt, u.UserName, u.Name, u.Email, ClassList(u.Class).Basic()}
}

type UserList []User

func (us UserList) Basic() []APIUserBasic {
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

func CreateUser(username, name, email string) (APIUserBasic, error) {
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return APIUserBasic{}, err
	// }
	user := User{
		UserName: username,
		Name:     name,
		Email:    email,
	}

	err := Create(&user)
	if err != nil {
		return APIUserBasic{}, err
	}
	return APIUserBasic{
		user.ID,
		user.UserName,
		user.Name,
		user.Email,
	}, nil
}

func GetUserByUserID(uid uint) (User, error) {
	user := User{}
	structQuery := User{ID: uid}
	// err := QueryOne(&structQuery, &user)
	// if err != nil {
	// 	return user, err
	// }
	// structQuery.ID = user.ID
	// err = db.Model(&structQuery).Association("Class").Find(&user.Class)
	err := db.Model(&structQuery).Preload("Class").Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	user := User{}
	structQuery := User{UserName: username}
	err := QueryOne(&structQuery, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func QueryUsersByUsername(usernames []string, data *[]User) error {
	err := db.Table("execode.users").Where("user_name IN ?", usernames).Find(data).Error
	return err
}
