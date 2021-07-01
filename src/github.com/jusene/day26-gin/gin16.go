package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type Person1 struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
	CreateTime time.Time `form:"createTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	route := gin.Default()
	route.GET("/testing", startPage1)
	route.Run()
}

func startPage1(ctx *gin.Context) {
	var person Person1
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	fmt.Println(ctx.ShouldBindQuery(&person))
	if ctx.ShouldBindQuery(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}
	ctx.String(200, "Success")
}
