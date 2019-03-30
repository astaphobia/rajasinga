package controllers

type (
	ErrorResponse interface {
		error
		Status() int
	}

	StatusError struct {
		Code int
		Err  error
	}
)

func (se StatusError) Error() string {
	return se.Err.Error()
}
func (se StatusError) Status() int {
	return se.Code
}
