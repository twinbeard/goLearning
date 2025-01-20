package initialize

import (
	"github.com/twinbeard/goLearning/global"
	"github.com/twinbeard/goLearning/internal/database"
	"github.com/twinbeard/goLearning/internal/service"
	"github.com/twinbeard/goLearning/internal/service/impl"
)

func InitServiceInterface() {
	// User service interface
	queries := database.New(global.Mdbc)
	service.InitUserLogin(impl.NewUserLoginImpl(queries))

}
