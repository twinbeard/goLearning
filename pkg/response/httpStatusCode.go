package response

const (
	ErrCodeSuccess    = 200 // Success
	ErrCodeBadRequest = 400 // Email is invalid
	ErrInvalidToken   = 401 // Invalid token
)

// message
var msg = map[int]string{
	ErrCodeSuccess:    "Success",
	ErrCodeBadRequest: "Email is invalid",
	ErrInvalidToken:   "Invalid token",
}
