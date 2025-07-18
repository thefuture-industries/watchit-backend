package httperr

func (e *ServerError) Error() string {
	return e.Msg
}

func (e *ServerError) StatusCode() int {
	return e.Code
}
