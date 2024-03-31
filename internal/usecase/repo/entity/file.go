package entity

type File struct {
	ID     int64  `db:"id" structs:"-" json:"id"`
	Bucket string `db:"bucket" structs:"bucket" json:"bucket"`
	Path   string `db:"path" structs:"path" json:"path"`
}
