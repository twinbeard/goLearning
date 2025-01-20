package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/twinbeard/goLearning/cmd/swag/docs"
	"github.com/twinbeard/goLearning/internal/initialize"
)

// @title           API DOCUMENT ECOMMERCE BACKEND SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/twinbeard/goLearning

// @contact.name   GO LEARNING
// @contact.url    github.com/twinbeard/goLearning
// @contact.email  truongvu.work@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8008
// @BasePath  /v1/2024
// @schema   http

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := initialize.Run()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8008")
}
