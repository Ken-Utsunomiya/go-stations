package model

type ErrNotFound struct {
	err error
}

func (e *ErrNotFound) Error() string {
	return e.err.Error()
}
