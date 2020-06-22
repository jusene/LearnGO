package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jusene/day27/docs"
	"github.com/jusene/day27/models"
	"net/http"
)

// @summary get path param
// @description swagger example get path param
// @tags get
// @accept json
// @produce json
// @param name path string true "name"
// @success 200 {object} models.Res
// @failure 404 {object} models.Err
// @router /get/{name} [get]
func GetPathParam(c *gin.Context) {
	name := c.Param("name")
	if name == "jusene" {
		c.JSON(http.StatusOK, models.Res{Name: name})
	} else {
		c.JSON(http.StatusNotFound, models.Err{Code: http.StatusNotFound, Msg: "not found"})
	}
}

// @summary get
// @description swagger example get
// @tags get
// @accept json
// @produce json
// @success 200 {object} models.Ress
// @router /get [get]
func Get(c *gin.Context) {
	user1 := models.Res{
		Name: "name",
		Msg:  "found",
	}
	user2 := models.Res{
		Name: "jusene",
		Msg:  "found",
	}
	c.JSON(http.StatusOK, models.Ress{Names: []models.Res{user1, user2}})
}

// @Summary 下载文件
// @Description
// @Tags file
// @Param filename query string true "file name"
// @Success 200 string ok
// @Router /download [get]
func Download(c *gin.Context) {
	filename := c.DefaultQuery("filename", "")
	// 对下载的文件重命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./files/" + filename)
}
