package postgres

const (
	UniqueConstraintError = "unique_constraint_error"
)

var SqlErrors = []string{
	UniqueConstraintError,
}
