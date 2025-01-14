package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/twinbeard/goLearning/pkg/logger"
	"github.com/twinbeard/goLearning/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
