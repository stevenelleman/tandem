package resp

import (
	"log"

	"github.com/gin-gonic/gin"
)

// TODO: Better error handling
func ReturnError(c *gin.Context, status int, err error) {
	log.Println("Error: %v", err)
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
