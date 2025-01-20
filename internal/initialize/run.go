package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/global"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	// cách biến của global.<variable> sẽ được sử dụng trong toàn bộ project
	LoadConfig() // Load configuration data from local.yaml and assign to global.Config
	fmt.Println("\033[32m CONFIGURATION LOADING SUCCESS\033[0m")
	InitLogger() // Initialize logger and assign to global.Logger variable
	fmt.Println("\033[32m LOGGER INITIALIZATION SUCCESS\033[0m")
	global.Logger.Info("Config Log okay!!", zap.String("OKIE DOKIE", "Sucess"))
	InitMysqlC()
	fmt.Println("\033[32m SQL DATABASE INITIALIZATION SUCCESS\033[0m")
	InitServiceInterface()
	fmt.Println("\033[32m SERIVCE INTERFACE INITIALIZATION SUCCESS\033[0m")
	InitRedis()
	fmt.Println("\033[32m REDIS INITIALIZATION SUCCESS\033[0m")
	InitKafka("localhost:19092", "otp-auth-topic")
	fmt.Println("\033[32m KAFKA INITIALIZATION SUCCESS\033[0m")
	r := InitRouter()
	fmt.Println("\033[32m ROUTER INITIALIZATION SUCCESS\033[0m")
	return r
}
