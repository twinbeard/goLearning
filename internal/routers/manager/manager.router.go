package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	/* Router của Superuser (admin)
	1. Đầu tiên admin sẽ đăng nhập vào localhost.com/admin/login để login in nên đây là public router
	2. Sau khi login thành công thì admin sẽ có thể access vào các router khác như localhost.com/admin/user/active_user để active hoặc deactive người bán hàng nên đây là private router

	*/

	// public router -> user ở đây nhé devs chứ không phải người quản lý đơn hàng nên phải login
	// Router.Group (quản lý router theo group)

	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	adminRouterPrivate.Use(
	// middlewares.Limiter(), // Limit request
	// middlewares.Authentication(), // Check user's token
	// middlewares.Permission(), // Check user's role permission to access this route

	)
	{
		adminRouterPrivate.GET("/active_user") // admin sẽ có quyền deactivate or activate người bán hàng và đơn hàng của người bán hàng
	}
}
