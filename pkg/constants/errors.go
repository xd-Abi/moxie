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
	ErrInternal              = &MoxieError{code: 500, message: "Internal error"}
	ErrEmailEmpty            = &MoxieError{code: 400, message: "Email can not be empty or whitespace"}
	ErrEmailInvalid          = &MoxieError{code: 400, message: "Email is not valid"}
	ErrEmailAlreadyExists    = &MoxieError{code: 400, message: "Email already exists"}
	ErrUsernameEmpty         = &MoxieError{code: 400, message: "Username can not be empty or whitespace"}
	ErrUsernameAlreadyExists = &MoxieError{code: 400, message: "Username already exists"}
	ErrPasswordEmpty         = &MoxieError{code: 400, message: "Password can not be empty or whitespace"}
	ErrJwtEmpty              = &MoxieError{code: 400, message: "Jwt can not be empty or whitespace"}
	ErrJwtInvalid            = &MoxieError{code: 400, message: "Jwt is invalid"}
	ErrSubjectEmpty          = &MoxieError{code: 400, message: "Subject can not be empty or whitespace"}
)
