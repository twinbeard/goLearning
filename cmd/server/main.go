package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/twinbeard/goLearning/cmd/swag/docs"
	"github.com/twinbeard/goLearning/internal/initialize"
)

// TEST ONLY - Nhớ xoá
var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count_total",
		Help: "Total number of ping requests.",
	},
)

// TEST ONLY - Nhớ xoá
func ping(c *gin.Context) {
	pingCounter.Inc() //1 2 3
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

// @title           API DOCUMENT ECOMMERCE BACKEND SHOPDEVGO
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/twinbeard/goLearning

// @contact.name   GO LEARNING
// @contact.url    github.com/twinbeard/goLearning
// @contact.email  truongvu.work@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema   http

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := initialize.Run()
	prometheus.MustRegister(pingCounter) // TEST ONLY - Nhớ xoá: Đăng ký metric
	r.GET("/ping/200", ping)             // TEST ONLY - Nhớ Xoá Cái này dùng để test graphana thôi .Nhớ xoá
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8002")
}
