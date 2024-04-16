package entity

import (
	"database/sql"
	"github.com/pets-shelters/backend-svc/pkg/date"
	"time"
)

type TaskExecution struct {
	ID     int64         `db:"id" structs:"-"`
	TaskID int64         `db:"task_id" structs:"task_id"`
	UserID sql.NullInt64 `db:"user_id" structs:"user_id"`
	Date   date.Date     `db:"date" structs:"date"`
	DoneAt time.Time     `db:"done_at" structs:"done_at"`
}

type TaskExecutionForList struct {
	UserID sql.NullInt64 `db:"user_id" structs:"user_id"`
	Date   date.Date     `db:"date" structs:"date"`
	DoneAt time.Time     `db:"done_at" structs:"done_at"`
}

type TaskExecutionForListNull struct {
	UserID sql.NullInt64 `db:"user_id" structs:"user_id"`
	Date   date.Date     `db:"date" structs:"date"`
	DoneAt sql.NullTime  `db:"done_at" structs:"done_at"`
}
