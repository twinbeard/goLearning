package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/routers"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine // declare gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode) // set GIN to "debug" mode
		gin.ForceConsoleColor()    // In ra màu sắc cho dễ nhìn ởconsole
		r = gin.Default()          // create Router Instance with default middleware (nó sẽ in lên console gin debug log từ bản thân chương trình)+ manual add middleware
	} else {
		gin.SetMode(gin.ReleaseMode) // set GIN to "release/production" mode
		r = gin.New()                // create Router Instance without default middleware (đồng thời bỏ chức debug logging của chương trình) and you need to manually add any middleware you want to use.
	}
	// middleware
	r.Use() // Logging
	r.Use() // Cross
	r.Use() // limiter global
	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("v1/2024")
	{
		MainGroup.GET("/checkStatus") // check status of server
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)

	}
	{
		managerRouter.InitAdminRouter(MainGroup)
		managerRouter.InitUserRouter(MainGroup)
	}
	return r
}
