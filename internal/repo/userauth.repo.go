package repo

import (
	"fmt"
	"time"

	"github.com/twinbeard/goLearning/global"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUserAuthRepository.
func (u *userAuthRepository) AddOTP(email string, otp int, expirationTime int64) error {
	// panic("unimplemented")
	key := fmt.Sprintf("user:%s:otp", email) // user:email:otp
	// Set the OTP and expiration time in Redis
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err() // SetEx sets the value and expiration of a key in Redis
	// Err() returns the error status of the command

}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
