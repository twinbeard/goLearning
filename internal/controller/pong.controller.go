package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// * PongController blueprint - Giống class trong OOP
type PongController struct {
}

// *	Giống constructor trong class
// NewPongController is a function that CREATE PongController instance and return its address
func NewPongController() *PongController {
	return &PongController{}
}

// * Giống method trong class
// uc is receiver which tell us the Pong function belongs to PongController
// Pong is a function of PongController. To use it, we have to create a PongController instance and call this function
func (uc *PongController) Pong(c *gin.Context) {
	// name := c.Param("name")        // Extract path params "name" => http://localhost:8080/v1/2024/ping/abc
	name := c.DefaultQuery("name", "Truong") // Extract path params "name" => http://localhost:8080/v1/2024/ping/
	uid := c.Query("uid")                    // Extract query params "uid" => http://localhost:8080/v1/2024/ping?uid=123
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
	})
}
