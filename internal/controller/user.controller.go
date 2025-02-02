package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/internal/service"
	"github.com/twinbeard/goLearning/internal/vo"
	"github.com/twinbeard/goLearning/pkg/response"
)

// Bình thường

// *Giống class trong OOP
// UserController blueprint
// type UserController struct {
// 	userService *service.UserService
// }

// *Giống constructor của clas nhưng đây cách viết khác
// NewUserController is a function that CREATE UserController instance and return its address
// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// * Giống method trong class
// uc is receiver which tell us the Pong function belongs to UserController
// User is a function of UserController. To use it, we have to create a UserController instance and call this function
// func (uc *UserController) GetUserByID(c *gin.Context) {
// name := c.Param("name")        // Extract path params "name" => http://localhost:8080/v1/2024/ping/abc
// name := c.DefaultQuery("name", "Truong") // Extract path params "name" => http://localhost:8080/v1/2024/ping/
// 	uid := c.Query("uid") // Extract query params "uid" => http://localhost:8080/v1/2024/ping?uid=123
// 	response.SuccessResponse(
// 		c, 2001, map[string]interface{}{
// 			"message": "pong" + uc.userService.GetUserByID(),
// 			"uid":     uid,
// 		},
// 	)
// }

// INTERFACE
type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamsInvald, err.Error())
		return
	}
	fmt.Printf("Email Params: %s\n", params)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}
