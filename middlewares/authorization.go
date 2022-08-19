package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func Authorization(redis_cache *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetHeaderAuthorization(c)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Authorization key in header not found"})
			c.Abort()
			return
		}

		// ส่ง Token ไปตรวจสอบว่าได้รับสิทธิ์เข้าใช้งานหรือไม่
		isToken, err := VerifyToken("accessToken",token, redis_cache)
		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message":"Authorization falil because of timeout."})
			c.Abort()
			return
		}

		if isToken {
			c.Next()
		}
	}

}


