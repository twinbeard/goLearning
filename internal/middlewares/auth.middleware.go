package middlewares

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/internal/utils/auth"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the request url path
		urlPath := c.Request.URL.Path
		log.Println("Request URL Path: ", urlPath)
		// Chech headers authorization
		jwtToken, valid := auth.ExtractBearerToken(c)
		if !valid {
			c.AbortWithStatusJSON(401, gin.H{
				"message": 40001,
				"err":     "Unauthorized",
				"detail":  "",
			})
			return
		}
		// validate jwt token by subject
		claims, err := auth.VerifyTokenSubject(jwtToken)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": 40001,
				"err":     "Invalid token",
				"detail":  "",
			})
			return
		}
		// update claims to context
		// => Nghĩa là chúng ta sẽ lưu thông tin của claim.subject (thông tin user) vào context. Từ đó, chúng ta có thể sử dụng thông tin này ở bất kỳ đâu trong hệ thống từ mỗi request của user. Đây lấy thông tin chúng ta dùng ctx.value("__")
		log.Println("Claims:::UUID ", claims.Subject) // 11cliToken
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}
