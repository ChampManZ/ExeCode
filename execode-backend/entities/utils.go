package entities

import (
	"time"

	"gorm.io/gorm"
)

type swaggoGorm struct {
	ID        int
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
}

func Paginate(pageSize int, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
