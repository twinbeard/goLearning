package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/twinbeard/goLearning/pkg/logger"
	"github.com/twinbeard/goLearning/setting"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB // GORM: Bỏ cái hàng này nhé vì mình sqlc thay cho gorm để thao tác với database
	Mdbc          *sql.DB  // SQLC
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
)
