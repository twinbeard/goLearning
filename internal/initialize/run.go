package initialize

import (
	"fmt"

	"github.com/twinbeard/goLearning/global"
	"go.uber.org/zap"
)

func Run() {
	// cách biến của global.<variable> sẽ được sử dụng trong toàn bộ project
	LoadConfig() // Load configuration data from local.yaml and assign to global.Config
	m := global.Config.Mysql
	fmt.Println("Loading configuration mySQL:", m.Username, m.Password, m.Host)
	InitLogger() // Initialize logger and assign to global.Logger variable
	global.Logger.Info("Config Log okay!!", zap.String("OKIE DOKIE", "Sucess"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run()
}
