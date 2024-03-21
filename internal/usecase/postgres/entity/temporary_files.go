package entity

import "time"

type TemporaryFile struct {
	ID        int64     `db:"id" structs:"-" json:"id"`
	FileID    int64     `db:"file_id" structs:"file_id" json:"file_id"`
	UserID    int64     `db:"user_id" structs:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" structs:"created_at" json:"created_at"`
}
