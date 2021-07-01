package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jusene/day27/models"
	"net/http"
)

// @Summary delete example
// @Description delete example
// @Tags delete
// @Produce json
// @Accept json
// @Security ApiKeyAuth
// @Param name path string true "name"
// @Success 200 {object} models.Res
// @Failure 404 {object} models.Err
// @Router /delete/{name} [delete]
func Delete(c *gin.Context) {
	name := c.Param("name")
	if name != "jusene" {
		c.JSON(http.StatusNotFound, models.Err{
			Code: http.StatusNotFound,
			Msg:  "not found",
		})
		return
	}
	c.JSON(http.StatusOK, models.Res{
		Name: "jusene",
		Msg:  "delete ok",
	})
}
// @Summary delete example
// @Description
// @Accept json
// @Produce json
// @Tags delete
// @Security ApiKeyAuth
// @Param id path string true "id"
// @Param name query string true "name"
// @Success 200 {object} models.Res
// @Router /deletequery/{id} [delete]
func DeleteQuery(c *gin.Context) {
	id := c.Param("id")
	name := c.DefaultQuery("name", "jusene")

	c.JSON(http.StatusOK, models.Res{
		Name: fmt.Sprintf("%s %s",id, name),
		Msg:  "delete ok",
	})
}