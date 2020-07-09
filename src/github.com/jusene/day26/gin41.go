package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"log"
	"net/http"
	"time"
)

// 限流器
// Token Bucket 漏斗桶
// 保证桶的请求均速流出，不关心流入的速度，只要桶溢出了就拒绝服务

var limiter = ratelimit.NewBucket(time.Second, 100)

func main() {
	r := gin.Default()
	r.GET("/test", tokenRateLimiter(), func(context *gin.Context) {
		context.JSON(200, true)
	})

	r.Run(":9090")
}

func tokenRateLimiter() gin.HandlerFunc {
	fmt.Println("token create rate: ", limiter.Rate())
	fmt.Println("availite token: ", limiter.Available())
	return func(context *gin.Context) {
		if limiter.TakeAvailable(1) == 0 {
			log.Printf("available token: %d", limiter.Available())
			context.AbortWithStatusJSON(http.StatusTooManyRequests, "too many request")
		} else {
			context.Writer.Header().Set("X-RateLimit-Remaining", fmt.Sprintf("%d", limiter.Available()))
			context.Writer.Header().Set("X-RateLimit-Limit", fmt.Sprintf("%d", limiter.Capacity()))
			context.Next()
		}
	}
}
