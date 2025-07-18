package httperr

type HTTPError interface {
	error
	StatusCode() int
}

type ServerError struct {
	Msg  string
	Code int
}
