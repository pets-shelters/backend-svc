package entity

import (
	"time"
)

type Shelter struct {
	ID          int64     `db:"id" structs:"-" json:"id"`
	Name        string    `db:"name" structs:"name" json:"name"`
	Logo        int64     `db:"logo" structs:"logo" json:"logo"`
	City        string    `db:"city" structs:"city" json:"city"`
	PhoneNumber string    `db:"phone_number" structs:"phone_number" json:"phone_number"`
	Instagram   string    `db:"instagram" structs:"instagram" json:"instagram"`
	Facebook    string    `db:"facebook" structs:"facebook" json:"facebook"`
	CreatedAt   time.Time `db:"created_at" structs:"created_at" json:"created_at"`
}
