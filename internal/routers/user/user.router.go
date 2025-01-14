package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	// use middleware on private router (/user/...)
	userRouterPrivate.Use(
	// Creates a gin router instance with default middleware: logger and recovery (crash-free) middleware
	// middlewares.Limiter(), // Limit request
	// middlewares.Authentication(), // Check user's token
	// middlewares.Permission(), // Check user's role permission to access this route

	)
	{
		userRouterPrivate.GET("/get_info")
	}
}
