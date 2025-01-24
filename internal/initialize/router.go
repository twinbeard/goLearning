package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/middlewares"
	"github.com/twinbeard/goLearning/internal/routers"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine // declare gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode) // set GIN to "debug" mode
		gin.ForceConsoleColor()    // In ra màu sắc cho dễ nhìn ở console
		r = gin.Default()          // create Router Instance with default middleware (nó sẽ in lên console gin debug log từ bản thân chương trình)+ manual add middleware
	} else {
		gin.SetMode(gin.ReleaseMode) // set GIN to "release/production" mode
		r = gin.New()                // create Router Instance without default middleware (đồng thời bỏ chức debug logging của chương trình) and you need to manually add any middleware you want to use.
	}
	// middleware
	r.Use()                                                  // Logging
	r.Use()                                                  // Cross
	r.Use(middlewares.NewRateLimitter().GlobalRateLimiter()) // limiter global - 100req/1s
	r.GET("/ping/100", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong 100"})
	})

	r.Use(middlewares.NewRateLimitter().PublicAPIRateLimiter()) // limiter global - 80req/1s
	r.GET("/ping/80", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong 80"})
	})

	r.Use(middlewares.NewRateLimitter().UserAndPrivateRateLimiter()) // limiter global - 50req/1s
	r.GET("/ping/50", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong 50"})
	})

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
