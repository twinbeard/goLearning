package service

import (
	"context"

	"github.com/twinbeard/goLearning/internal/models"
)

// Mục tiêu interface là để dùng cho leader viết ra những yêu cầu cho những lập trình viên khác thực hiện
// File interface này chỉ dành cho LEADER viết
type (
	// ... other interfaces

	IUserLogin interface {
		Login(ctx context.Context, in *models.LoginInput) (codeResult int, out models.LoginOutput, err error)
		Register(ctx context.Context, in *models.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *models.VerifyInput) (out models.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (useId int, err error)
		// two-factor authentication
		IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error)
		// setup authentication
		SetupTwoFactorAuth(ctx context.Context, in *models.SetupTwoFactorAuthInput) (codeResult int, err error)
		// Verify Two Factor Authentication
		VerifyTwoFactorAuth(ctx context.Context, in *models.TwoFactorVerificationInput) (codeResult int, err error)
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
	localUserLogin IUserLogin
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("implement localUserAdmin not found for interface IUserAdmin")
	}
	return localUserAdmin
}
func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("implement localUserInfo not found for interface IUserInfo")
	}
	return localUserInfo
}
func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}
func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}
func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
