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
	ErrInternal               = &MoxieError{code: 500, message: "Internal error"}
	ErrUnauthorized           = &MoxieError{code: 400, message: "Unauthorized, Forbidden resource"}
	ErrUserNotFound           = &MoxieError{code: 400, message: "User not found"}
	ErrProfileNotFound        = &MoxieError{code: 400, message: "Profile not found"}
	ErrEmailEmpty             = &MoxieError{code: 400, message: "Email can not be empty or whitespace"}
	ErrEmailInvalid           = &MoxieError{code: 400, message: "Email is not valid"}
	ErrEmailAlreadyExists     = &MoxieError{code: 400, message: "Email already exists"}
	ErrUsernameEmpty          = &MoxieError{code: 400, message: "Username can not be empty or whitespace"}
	ErrUsernameAlreadyExists  = &MoxieError{code: 400, message: "Username already exists"}
	ErrUserIdEmpty            = &MoxieError{code: 400, message: "User id can not be empty or whitespace"}
	ErrPasswordEmpty          = &MoxieError{code: 400, message: "Password can not be empty or whitespace"}
	ErrPasswordInvalid        = &MoxieError{code: 400, message: "Password does not match"}
	ErrCurrentPasswordEmpty   = &MoxieError{code: 400, message: "Current password can not be empty or whitespace"}
	ErrCurrentPasswordInvalid = &MoxieError{code: 400, message: "Current password does not match"}
	ErrNewPasswordEmpty       = &MoxieError{code: 400, message: "New password can not be empty or whitespace"}
	ErrJwtEmpty               = &MoxieError{code: 400, message: "Jwt can not be empty or whitespace"}
	ErrJwtInvalid             = &MoxieError{code: 400, message: "Jwt is invalid"}
	ErrRefreshTokenEmpty      = &MoxieError{code: 400, message: "Refresh token is empty"}
	ErrRefreshTokenInvalid    = &MoxieError{code: 400, message: "Refresh token is invalid"}
	ErrSubjectEmpty           = &MoxieError{code: 400, message: "Subject can not be empty or whitespace"}
)
