package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var secrets = gin.H{
	"foo": gin.H{"email": "foo@bar.com", "phone": "123433"},
}

func main() {
	r := gin.Default()

	authorizd := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	}))

	authorizd.GET("/secrets", func(context *gin.Context) {
		user := context.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET"})
		}
	})

	r.Run()
}
