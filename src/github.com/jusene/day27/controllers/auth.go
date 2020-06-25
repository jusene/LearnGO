package controllers

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) bool {
	if c.GetHeader("Authorization") != "1234" {
		return false
	}
	return true
}
