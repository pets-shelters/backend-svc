package entity

import (
	"time"
)

type Shelter struct {
	ID          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Logo        string    `db:"logo" json:"logo"`
	City        string    `db:"city" json:"city"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Instagram   string    `db:"instagram" json:"instagram"`
	Facebook    string    `db:"facebook" json:"facebook"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
