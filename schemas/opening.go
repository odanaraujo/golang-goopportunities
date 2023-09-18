package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Link     string
	Remote   bool
	Salary   float64
}

type OpeningResponse struct {
	ID       uint      `json:"id"`
	Role     string    `json:"role"`
	Company  string    `json:"company"`
	Location string    `json:"location"`
	Link     string    `json:"link"`
	Remote   bool      `json:"remote"`
	Salary   float64   `json:"salary"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at,omitempty"`
}
