package entity

import "time"

type TemporaryFile struct {
	ID        int64     `db:"id" structs:"-"`
	FileID    int64     `db:"file_id" structs:"file_id"`
	UserID    int64     `db:"user_id" structs:"user_id"`
	CreatedAt time.Time `db:"created_at" structs:"created_at"`
}
