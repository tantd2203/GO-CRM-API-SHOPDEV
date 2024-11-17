package response

const (
	ErrCodeSuccess      = 201 // success
	ErrCodeParamInvalid = 203 // email invalid
	ErrInvalidToken     = 301
)

// message
var msg = map[int]string{

	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "ErrInvalidToken",
}
