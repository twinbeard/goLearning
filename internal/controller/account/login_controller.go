package account

import (
	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/models"
	"github.com/twinbeard/goLearning/internal/service"
	"github.com/twinbeard/goLearning/pkg/response"
	"go.uber.org/zap"
)

type cUserLogin struct{}

// Login is a instance of cUserLogin
var Login = new(cUserLogin)

// User Login
// @Summary      User Login
// @Description  User Login
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body models.LoginInput true "payload"
// @Failure      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/login    [post]
func (c *cUserLogin) Login(ctx *gin.Context) {
	// Implement logic for login
	var params models.LoginInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}
	codeResult, dataRs, err := service.UserLogin().Login(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeResult, dataRs)
}

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered, send OTP to user's email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body models.RegisterInput true "payload"
// @Failure      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register     [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var params models.RegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {

		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}
	codeStatus, err := service.UserLogin().Register(ctx, &params)
	if err != nil {
		global.Logger.Error("Error registering user OTP : ", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}

// Verify OTP Login By User
// @Summary      Verify OTP Login By User
// @Description  Verify OTP Login By User
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body models.VerifyInput true "payload"
// @Failure      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/verify_account    [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params models.VerifyInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}
	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}

// Update Password Register
// @Summary      Update Password Register
// @Description  Update Password Register
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body models.UpdatePasswordRegisterInput true "payload"
// @Failure      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/update_pass_register    [post]
func (c *cUserLogin) UpdatePasswordRegister(ctx *gin.Context) {
	var params models.UpdatePasswordRegisterInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}
	result, err := service.UserLogin().UpdatePasswordRegister(ctx, params.UserToken, params.UserPassword)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamsInvald, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
