package util

type Error struct {
	msg string
}

func (error *Error) Error() string {
	return error.msg
}

func ErrInvalidPath() *Error {
	return &Error{"The first argument should be a valid path"}
}

func ErrInvalidImportFile() *Error {
	return &Error{"Provided file is not a valid file to import"}
}
