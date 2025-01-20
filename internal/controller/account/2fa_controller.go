package account

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/internal/models"
	"github.com/twinbeard/goLearning/internal/service"
	"github.com/twinbeard/goLearning/internal/utils/context"
	"github.com/twinbeard/goLearning/pkg/response"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct {
}

// User Set Up Two Factor Authentication
// @Summary      User Set Up Two Factor Authentication
// @Description  User Set Up Two Factor Authentication
// @Tags         account two-factor authentication
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Authorization token"
// @Param        payload body models.SetupTwoFactorAuthInput true "payload"
// @Failure      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/setup   [post]
func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params models.SetupTwoFactorAuthInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid set")
		return
	}
	// get UserId from uuid (token) -  Get userId from context
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid set")
		return
	}
	log.Println("userId: ", userId)
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}

// User Verify Two Factor Authentication
// @Summary      User Verify Two Factor Authentication
// @Description  User Verify Two Factor Authentication
// @Tags         account two-factor authentication
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Authorization token"
// @Param        payload body models.TwoFactorVerificationInput true "payload"
// @Failure      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/verify   [post]
func (c *sUser2FA) VerifyTwoFactorAuth(ctx *gin.Context) {
	var params models.TwoFactorVerificationInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, "Missing or invalid set")
		return
	}
	// get UserId from uuid (token) -  Get userId from context
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, "User ID is invalid")
		return
	}
	log.Println("userId: VerifyTwoFactorAuth: ", userId)
	params.UserId = uint32(userId)
	codeResult, err := service.UserLogin().VerifyTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthVerifyFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}
