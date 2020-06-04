package handlers

import (
	"github.com/gin-gonic/gin"
	urlschema "github.com/gorilla/schema"
)

var (
	// Terrible to have global variable, but otherwise need to create on per-request basis
	urlDecoder = urlschema.NewDecoder()
)

// TODO: Apply to POST/PUT requests
func MapParams(c *gin.Context, dst interface{}) error {
	v := c.Request.URL.Query()
	err := urlDecoder.Decode(dst, v)
	if err != nil {
		return err
	}
	return nil
}
