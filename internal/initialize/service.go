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
	// create
	userLoginImplInstance := impl.NewUserLoginImpl(queries) // Create instance of a struct to implement interace (User Login Struct)
	service.InitUserLogin(userLoginImplInstance)            // Assign instance of UserLoginImpl to UserLogin interface

}
