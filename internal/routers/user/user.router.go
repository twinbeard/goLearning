package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/internal/controller/account"
	"github.com/twinbeard/goLearning/internal/middlewares"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	fmt.Println("Init User Router")
	// public router
	/* Non-dependency injection
	Non-dependency => Có nghĩa là nếu có các function phụ thuộc vào nhau mà chúng ta không group lại thì sau này khi codespace nhiều lên
	Dependency => Có nghĩa là group các hàm phụ thuộc vào nhau lại
	sẽ khó sửa nên bây giờ chúng ta sẽ dùng PACKAGE wire đến group các hàm phụ thuộc vào nhau rồi sau này chỉ cần gọi 1 hàm là được
	install wire: go get github.com/google/wire/cmd/wire
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userController:= controller.NewUserController(us)
	userRouterPublic := Router.Group("/user", userHandlerNonDependency.Register)
	*/
	//* Dependency injection (DI) => Dùng package WIRE để inject dependency
	// userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := Router.Group("/user")
	{
		// userRouterPublic.POST("/register", userController.Register)
		// userRouterPublic.POST("/otp")
		userRouterPublic.POST("/register", account.Login.Register)
		userRouterPublic.POST("/update_pass_register", account.Login.UpdatePasswordRegister)
		userRouterPublic.POST("/verify_account", account.Login.VerifyOTP)
		userRouterPublic.POST("/login", account.Login.Login) // Dùng cách  viết interface mới
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// use middleware on private router (/user/...)
	userRouterPrivate.Use(
		// Creates a gin router instance with default middleware: logger and recovery (crash-free) middleware
		middlewares.AuthenMiddleware(),
	// middlewares.Limiter(), // Limit request
	// middlewares.Authentication(), // Check user's token
	// middlewares.Permission(), // Check user's role permission to access this route

	)
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.POST("/two-factor/setup", account.TwoFA.SetupTwoFactorAuth)
	}
}
