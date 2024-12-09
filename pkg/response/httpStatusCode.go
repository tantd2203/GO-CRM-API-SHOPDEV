package response

const (
	ErrCodeSuccess      = 201 // success
	ErrCodeParamInvalid = 203 // email invalid
	ErrInvalidToken     = 301

	// Register code

	ErrCodeUserHasExists = 501
)

// message
var msg = map[int]string{

	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "ErrInvalidToken",
	ErrCodeUserHasExists: "user has already registered",
}
