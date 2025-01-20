package impl

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/consts"
	"github.com/twinbeard/goLearning/internal/database"
	"github.com/twinbeard/goLearning/internal/models"
	"github.com/twinbeard/goLearning/internal/utils"
	"github.com/twinbeard/goLearning/internal/utils/auth"
	"github.com/twinbeard/goLearning/internal/utils/crypto"
	"github.com/twinbeard/goLearning/internal/utils/random"
	"github.com/twinbeard/goLearning/internal/utils/sendto"
	"github.com/twinbeard/goLearning/pkg/response"
)

// THIS IS FILE TO IMPLEMENT INTERFACE IUserLogin

type sUserLogin struct {
	// Implement the IUserLogin interface here
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

//* ------------------TWO-FACTOR AUTHENTICATION ---------------------------------------------------

func (s *sUserLogin) IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error) {
	return 200, true, nil
}

// setup authentication
func (s *sUserLogin) SetupTwoFactorAuth(ctx context.Context, in *models.SetupTwoFactorAuthInput) (codeResult int, err error) {
	//1. Check if user has enabled two-factor authentication
	isTwoFactorEnable, err := s.r.IsTwoFactorEnabled(ctx, in.UserId)
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, err
	}
	//1.1 If user has enabled two-factor authentication then return true
	if isTwoFactorEnable > 0 {
		return response.ErrCodeTwoFactorAuthSetupFailed, fmt.Errorf("two-factor authentication has already enabled")
	}
	//2.If user has not enabled two-factor authentication then continue to setup new type of two-factor authentication such as email or mobile
	err = s.r.EnableTwoFactorTypeEmail(ctx, database.EnableTwoFactorTypeEmailParams{
		UserID:            in.UserId,
		TwoFactorAuthType: database.PreGoAccUserTwoFactor9999TwoFactorAuthTypeSMS,
		TwoFactorEmail:    sql.NullString{String: in.TwoFactorEmail, Valid: true},
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthSetupFailed, fmt.Errorf("two-factor authentication has already enabled")
	}
	//3. If user enable two-factor authentication successfully then send OTP to user in.TwoFactorEmail
	keyUserTwoFactor := crypto.GetHash("2fa" + strconv.Itoa(int(in.UserId)))
	go global.Rdb.Set(ctx, keyUserTwoFactor, "123456", time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()

	return response.ErrCodeSuccess, nil
}

// Verify Two Factor Authentication
func (s *sUserLogin) VerifyTwoFactorAuth(ctx context.Context, in *models.TwoFactorVerificationInput) (codeResult int, err error) {
	return 200, nil
}

//* ------------------ END TWO-FACTOR AUTHENTICATION ---------------------------------------------------

func (s *sUserLogin) Login(ctx context.Context, in *models.LoginInput) (codeResult int, out models.LoginOutput, err error) {
	/*
		1. Khi login sẽ sinh ra access token và refresh token
		Sau đó, sẽ sinh ra 1 cái token để sử dụng trong các service
	*/

	userBase, err := s.r.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// 1. Check if password is matched with the password in database
	if !crypto.MatchingPassword(userBase.UserPassword, in.UserPassword, userBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("password does not match")
	}
	// 2. Check two-factor authentication
	// 2.1 Check if user has enabled two-factor authentication

	//

	// 3. Generate access token and refresh token
	go s.r.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: in.UserPassword, // Không cần passworrk
	}) // Đây là goroutine dùng để lưu thông tin login vào bảng login_user_base trong mysql ở background nên không cần quan tâm đến kết quả trả về
	// Chính vì thế chúng tao không cần wait goroutine

	// 5. Create UUID User
	subToken := utils.GenerateCliTokenUUID(int(userBase.UserID))
	log.Println("subToken", subToken)
	// 6. get user_info table
	infoUser, err := s.r.GetUser(ctx, uint64(userBase.UserID))

	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}
	// convert to json
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Error convert to json: %v", err)
	}
	// 7. Save infoUserJson to redis with key = subToken
	err = global.Rdb.SetEx(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err() // THời gian hết hạn = time của token
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("Error convert to json: %v", err)
	}
	// 8. Create JWT token
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return
	}
	return 200, out, nil
}

func (s *sUserLogin) Register(ctx context.Context, in *models.RegisterInput) (codeResult int, err error) {
	// Register implements IUserService.

	// 1. hashEmail // VerifyKey là email người dùng nhập vào ở front end
	fmt.Printf("Verify key is :::%s\n", in.VerifyKey)
	fmt.Printf("Verify type is :::%d\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("Hash key is :::%s\n", hashKey)

	// 2. check email exists in db
	userfound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExisted, err
	}
	if userfound > 0 {
		return response.ErrCodeUserHasExisted, fmt.Errorf("User has already registered")
	}

	// 3. Check if key exists in redis
	userKey := utils.GetUserKey(hashKey) //
	otp, err := global.Rdb.Get(ctx, userKey).Result()
	// Đưa cái này vào file utils
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed::", err)
		return response.ErrInvalidOTP, err
	case otp != "":
		return response.ErrCodeOtpCodeNotExists, fmt.Errorf("OTP code is not exists")
	}
	//4. Generate New OTP
	otpNew := random.GenerateSixDigitOtp()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("Otp is :::%d\n", otpNew)
	//5. Save OTP to redis
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	//6. Send OTP to user
	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOtp([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		// Save otp to mysql
		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
		})

		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		// 8. getlasId
		lastIdVerifyUser, err := result.LastInsertId()

		if err != nil {
			return response.ErrSendEmailOTP, err
		}
		log.Println("lastIdVerifyUser", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil
	case consts.MOBILE:
		return response.ErrCodeSuccess, nil

	}
	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *models.VerifyInput) (out models.VerifyOTPOutput, err error) {
	// 1. Check if hash key exists in redis or not
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("Hash key is :::%s\n", hashKey)
	optFound, err := global.Rdb.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}
	// VerifyCode là mã OTP người dùng nhập vào ở front end sau khi nhận được mã OTP qua email
	if in.VerifyCode != optFound {
		// Nếu như sai 3 lần trong 1 phút: Chưa làm nhé
		return out, fmt.Errorf("OTP code is not correct")
	}
	// 2. Check if hashkey of Email exists in mysql
	infoOTP, err := s.r.GetInfoOTP(ctx, in.VerifyKey)
	if err != nil {
		return out, err
	}
	// if hashkey of email exist in mysql then update status of email
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey) // is_verified = 1
	if err != nil {
		return out, err
	}
	// Use haskey as temporary token so that when user update information in front end then they will send this token to server for use to update in verify table also
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "Verify success"

	return out, err
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userID int, err error) {
	// Check if token exists in mysql table "user_verify"
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	// Check token is verified or not
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeOtpCodeNotExists, fmt.Errorf("OTP is not verified")
	}
	// If token is verified then check if token exists in redis or not
	// Update user_base table trong mysql
	userBase := database.AddUserBaseParams{}

	userBase.UserAccount = infoOTP.VerifyKey
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	userBase.UserSalt = userSalt

	userBase.UserPassword = crypto.HashPassword(password, userSalt)

	// add userBase to user_base table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	// add user_id to user_info table
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true}, // nickname = email
		UserAvatar:           sql.NullString{String: "", Valid: true},                // avatar = ""
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})
	user_id, err = newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	return int(user_id), nil
}
