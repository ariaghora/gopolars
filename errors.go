package gopolars

type PolarsError struct {
	message string
}

func (e *PolarsError) Error() string {
	return e.message
}
