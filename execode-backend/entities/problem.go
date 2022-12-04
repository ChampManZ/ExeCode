package entities

import (
	"time"

	"gorm.io/gorm"
)

type Problem struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index" swaggertype:"string"`
	ClassID        int            `gorm:"not null;uniqueIndex:idx_class_problem_name"`
	Class          Class
	ProblemName    string `gorm:"size:80;not null;uniqueIndex:idx_class_problem_name"`
	ProblemContent ProblemContent
}

type ProblemContent struct {
	ProblemID uint       `gorm:"primaryKey" json:"-"`
	Content   string     `gorm:"text;not null" json:"statement"`
	TestCases []TestCase `gorm:"foreignKey:ProblemID;references:ProblemID" json:"testcases"`
}

type TestCase struct {
	ProblemID  uint   `gorm:"primaryKey" json:"-"`
	TestCaseID uint   `gorm:"primaryKey" json:"-"`
	Input      string `gorm:"type:text;not null" json:"input"`
	Output     string `gorm:"type:text;not null" json:"output"`
}

type APIProblemBasic struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ProblemName string    `json:"problem_name"`
}

type APIProblemAdvanced struct {
	ID             uint           `json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	ProblemName    string         `json:"problem_name"`
	ProblemContent ProblemContent `json:"content"`
}

func (p Problem) Advanced() APIProblemAdvanced {
	return APIProblemAdvanced{
		ID:             p.ID,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		ProblemName:    p.ProblemName,
		ProblemContent: p.ProblemContent,
	}
}

func (p Problem) Basic() APIProblemBasic {
	return APIProblemBasic{
		p.ID,
		p.CreatedAt,
		p.UpdatedAt,
		p.ProblemName,
	}
}

type ProblemList []Problem

func (ps ProblemList) Basic() []APIProblemBasic {
	ret := make([]APIProblemBasic, len(ps))
	for i, u := range ps {
		ret[i] = u.Basic()
	}
	return ret
}

func GetProblemByID(id uint) (Problem, error) {
	query := Problem{ID: id}
	err := db.Preload("ProblemContent").Preload("ProblemContent.TestCases").Where(&query).First(&query).Error
	return query, err
}

func DeleteProblemByID(id uint) error {
	if err := db.Unscoped().Delete(&Problem{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func (p *Problem) BeforeDelete(tx *gorm.DB) (err error) {
	if err = tx.Delete(&TestCase{ProblemID: p.ID}).Error; err != nil {
		return
	}

	err = tx.Delete(&ProblemContent{ProblemID: p.ID}).Error
	return
}

func GetClassProblemsByID(cid uint) (problems []Problem, err error) {
	class := Class{ID: cid}
	err = db.Model(&class).Association("Lectures").Find(&problems)
	return
}
