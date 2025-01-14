package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// No public router for user (user ở đây là người quản lý đơn hàng của shop chứ 0 phải devs nhé)

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	// use middleware on private router (/user/...)
	userRouterPrivate.Use(
	// middlewares.Limiter(), // Limit request
	// middlewares.Authentication(), // Check user's token
	// middlewares.Permission(), // Check user's role permission to access this route

	)
	{
		userRouterPrivate.POST("/active_info") // admin sẽ có quyền deactivate or activate đơn hàng
	}
}
