package initialize

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/twinbeard/goLearning/global"
	"go.uber.org/zap"
)

var ctx = context.Background() // context is used to manage the lifecycle of the Redis connection

func InitRedis() {
	// Create a new Redis client
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password, // no password set
		DB:       r.Database, // use default DB
		PoolSize: 10,         // max number of connections in the pool
	})

	// Check if the connection is successful
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization Error:", zap.Error(err))
	}
	fmt.Println("Redis initialization successful")
	global.Rdb = rdb
	redisExample()
}

// redisExample is an example of using Redis to store data and retrieve data
func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err))
		return
	}

	global.Logger.Info("value score is::", zap.String("score", value))
}
