package handlers

import (
	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

func ReturnJSON(c *gin.Context, status int, obj interface{}) {
	c.JSON(status, gin.H{
		"obj": obj,
	})
}

func ReturnJSONList(c *gin.Context, status int, obj interface{}) {
	c.JSON(status, gin.H{
		"list": obj,
	})
}
