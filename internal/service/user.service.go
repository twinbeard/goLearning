package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/repo"
	"github.com/twinbeard/goLearning/internal/utils/crypto"
	"github.com/twinbeard/goLearning/internal/utils/random"
	"github.com/twinbeard/goLearning/pkg/response"
)

/* Cách viết thông thường cho service without interface

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUserByID() string {
	return us.userRepo.GetUserByID()
}

*/

// WITH INTERFACE
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	// Trong một user service có nhiều repo khác nhau => Nếu có nữa thì add vào đây
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// Register implements IUserService.

	// 0. hashEmail
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash email is :::%s\n", hashEmail)

	// 5. check OTP is available

	// 6. user spam ...

	// 1. check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		fmt.Printf("Email %s has existed\n", email)
		return response.ErrCodeUserHasExisted
	}

	// 2. new OTP ->
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("Otp is :::%d\n", otp)

	// 3. save OTP in Redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	fmt.Printf("Error is :::%v\n", err)
	if err != nil {
		return response.ErrInvalidOTP
	}
	//* Có 3 cách để OTP được gửi đến user => Cách 3 là tốt nhất
	// C1: Send OTP to user via email by Golang
	// 4. send Email OTP in text format
	// err = sendto.SendTextEmailOtp([]string{email}, "tlttmt@gmail.com", strconv.Itoa(otp))

	// 4. send Email OTP in HTML format
	// err = sendto.SendTemplateEmailOtp([]string{email}, "tlttmt@gmail.com", "auth-otp.html", map[string]interface{}{"otp": strconv.Itoa(otp)})
	// if err != nil {
	// 	return response.ErrSendEmailOTP
	// }
	/* C2: Send OTP task to java then java will send otp to user via email
	User Send email info to Golang then Golang send that to Java then Java send OTP back user via email

	*/
	// err = sendto.SendEmailToJavaByAPI(strconv.Itoa(otp), email, purpose)
	// fmt.Println("Send email by Java")
	// if err != nil {
	// 	return response.ErrSendEmailOTP
	// }
	/* C3 Send OTP task to kafka and java pick up that task from kafka and response otp via email to user  */
	body := make(map[string]interface{})
	body["otp"] = otp
	body["email"] = email

	bodyRequest, _ := json.Marshal(body)

	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: []byte(bodyRequest),
		Time:  time.Now(),
	}

	err = global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		return response.ErrSendEmailOTP
	}
	return response.ErrCodeSuccess
}

func NewUserService(userRepo repo.IUserRepository, useAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: useAuthRepo,
	}
}
