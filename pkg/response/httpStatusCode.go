package response

const (
	ErrCodeSuccess      = 201 // succes
	ErrCodeParamInvalid = 203 // email invalid

)

// message
var msg = map[int]string{

	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
}
