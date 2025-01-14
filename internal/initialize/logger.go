package initialize

import (
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
