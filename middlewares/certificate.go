package middlewares

import (
	"RU-Smart-Workspace/ru-smart-api/handlers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Certificate(redis_cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetHeaderAuthorization(c)
		fmt.Println(token)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Certificate key in header not found"})
			c.Abort()
			return
		}

		// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
		isToken, err := VerifyCertificateToken("accessToken", token, redis_cache)
		if err != nil {
			c.Error(err)
			c.Set("line", handlers.GetLineNumber())
			c.Set("file", handlers.GetFileName())
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Certificate falil because of timeout..."})
			c.Abort()
			return
		}

		if isToken {
			c.Next()
		}
	}

}
