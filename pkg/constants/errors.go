package constants

type MoxieError struct {
	message string
	// The code is based on default http status code
	code uint
}

func (e *MoxieError) Error() string {
	return e.message
}

func (e *MoxieError) Code() uint {
	return e.code
}

var (
	ErrInternal      = &MoxieError{code: 500, message: "Internal error"}
	ErrEmailEmpty    = &MoxieError{code: 400, message: "Email can not be empty or whitespace"}
	ErrEmailInvalid  = &MoxieError{code: 400, message: "Email is not valid"}
	ErrUsernameEmpty = &MoxieError{code: 400, message: "Username can not be empty or whitespace"}
)
