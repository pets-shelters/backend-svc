package entity

type File struct {
	ID     int64  `db:"id" structs:"-"`
	Bucket string `db:"bucket" structs:"bucket"`
	Path   string `db:"path" structs:"path"`
}
