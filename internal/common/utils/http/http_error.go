package http

func NewError(etype, message string, code int, data interface{}) Error {
	return Error{
		Message:   message,
		ErrorType: etype,
		Data:      data,
		Code:      code,
	}
}

type Error struct {
	Code      int
	ErrorType string
	Message   string
	Data      interface{}
}

func (e Error) Error() string {
	return e.Message
}
