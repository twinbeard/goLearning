package account

import (
	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/internal/models"
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
func (c sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params models.SetupTwoFactorAuthInput

	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid set")
		return
	}

	// get UserId from uuid (token)
}
