package model

// ErrNotFound expresses a custome error for 404
type ErrNotFound struct{}

func (e *ErrNotFound) Error() string {
	return "Not Found"
}
