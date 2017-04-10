package core

//Error struct used to report model errors
type Error struct {
	err    string
	Errors map[string]string
}

//NewModelError creates a new error model
func NewModelError(s string, errors map[string]string) Error {
	return Error{err: s, Errors: errors}
}

func (e Error) Error() string {
	return e.err
}
