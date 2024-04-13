package entity

type TaskAnimal struct {
	ID       int64 `db:"id" structs:"-"`
	AnimalID int64 `db:"animal_id" structs:"animal_id"`
	TaskID   int64 `db:"task_id" structs:"task_id"`
}
