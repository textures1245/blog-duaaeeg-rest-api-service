package entity

type CError struct {
	StatusCode int   `json:"status_code"`
	Err        error `json:"error"`
}

func (e *CError) Error() string {
	return e.Err.Error()
}

type ErrorHandler interface {
	PrismaAuthHandle(err CError) *CError
	PrismaPostHandle(err CError) *CError
}
