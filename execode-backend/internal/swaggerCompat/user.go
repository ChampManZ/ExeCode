package swaggercompat

import (
	"time"
)

type Response struct {
	Result interface{} `json:"result"`
}

type UserAdvanceWithRelation struct {
	ID        uint         `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	UserName  string       `json:"user_name"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Email     string       `json:"email"`
	Class     []ClassBasic `json:"classes"`
} // @name UserAdvanced
