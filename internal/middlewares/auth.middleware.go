package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/pkg/response"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(c, response.ErrInvalidToken, "")
			c.Abort() // Stop the middleware chain
			return
		}
		c.Next() // Call the next middleware

	}
}
