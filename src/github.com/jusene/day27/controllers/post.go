package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jusene/day27/models"
	"net/http"
)

// @summary post path param
// @description swagger example post path param
// @tags post
// @accept json
// @produce json
// @param name path string true "name"
// @success 200 {object} models.Res
// @failure 404 {object} models.Err
// @router /post/{name} [post]
func PostPathParam(c *gin.Context) {
	name := c.Param("name")
	if name == "jusene" {
		c.JSON(http.StatusOK, models.Res{Name: name})
	} else {
		c.JSON(http.StatusNotFound, models.Err{Code: http.StatusNotFound, Msg: "not found"})
	}
}

// @summary post body
// @description swagger example post body
// @tags post
// @accept json
// @produce json
// @param user body models.User true "name"
// @success 200 {object} models.User
// @failure 400 {object} models.Err
// @router /post [post]
func Post(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Msg:  "request error",
		})
		return
	}
	if len(user.Hobbys) != 0 {
		for _, h := range user.Hobbys {
			fmt.Println(h.Name)
		}
	}
	c.JSON(http.StatusOK, models.User{
		Name: user.Name,
		Age:  user.Age,
		// Hobbys: []models.Hobby{models.Hobby{Name: "ride"}, models.Hobby{Name: "car"}},
	})
}

// @summary post header
// @description swagger example post header
// @tags post
// @accept json
// @produce json
// @tags post
// @security ApiKeyAuth
// @param user body models.User true "name"
// @success 200 {object} models.User
// @failure 400 {object} models.Err
// @failure 401 {object} models.Err
// @router /postheader [post]
func PostHeader(c *gin.Context) {
	var user models.User
	if c.GetHeader("Authorization") != "1234" {
		c.JSON(http.StatusUnauthorized, models.Err{
			Code: http.StatusUnauthorized,
			Msg:  "no authorized",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Err{
			Code: http.StatusBadRequest,
			Msg:  "request error",
		})
		return
	}
	if len(user.Hobbys) != 0 {
		for _, h := range user.Hobbys {
			fmt.Println(h.Name)
		}
	}
	c.JSON(http.StatusOK, models.User{
		Name: user.Name,
		Age:  user.Age,
		// Hobbys: []models.Hobby{models.Hobby{Name: "ride"}, models.Hobby{Name: "car"}},
	})
}
