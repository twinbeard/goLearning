package response

const (
	ErrCodeSuccess          = 20001 // Success
	ErrCodeBadRequest       = 20003 // Email is invalid
	ErrInvalidToken         = 30001 // Invalid token
	ErrInvalidOTP           = 30002 // Invalid OTP
	ErrSendEmailOTP         = 30003 // Send email OTP failed
	ErrCodeParamsInvald     = 40001 // Params invalid
	ErrCodeUserHasExisted   = 50001 // User has already registered
	ErrCodeOtpCodeNotExists = 60009
	ErrCodeUserOtpNotExists = 60008
	//User authentication
	ErrCodeAuthFailed = 40005
	// Two-factor-authentication
	ErrCodeTwoFactorAuthSetupFailed = 80001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:          "Success",
	ErrCodeBadRequest:       "Email is invalid",
	ErrInvalidToken:         "Invalid token",
	ErrInvalidOTP:           "Invalid OTP",
	ErrSendEmailOTP:         "Send email OTP failed",
	ErrCodeParamsInvald:     "Params invalid",
	ErrCodeUserHasExisted:   "User has already registered",
	ErrCodeOtpCodeNotExists: "OTP code exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrCodeAuthFailed:       "User authentication failed",
	// Two-factor-authentication
	ErrCodeTwoFactorAuthSetupFailed: "Two-factor-authentication setup failed",
}
