package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
)

// 限流器
// Leaky Bucket
// 漏斗桶  限制数据的平均传输速度，比如限制发帖频率等
/*
func main() {
	r1 := ratelimit.New(100)
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := r1.Take()
		fmt.Println(i, now.Sub(prev))
		prev = now
	}
}
*/

var r1 ratelimit.Limiter

func leakyBucketRateLimiter() gin.HandlerFunc {
	//prev := time.Now()
	return func(context *gin.Context) {
		r1.Take()
		//fmt.Println(now.Sub(prev)) // 打印下多次请求的时间间隔
		//prev = now
	}
}

func main() {
	r := gin.Default()
	r.GET("/test", leakyBucketRateLimiter(), func(context *gin.Context) {
		context.JSON(200, true)
	})
	r.Run(":9091")
}

func init() {
	r1 = ratelimit.New(1000)
}
