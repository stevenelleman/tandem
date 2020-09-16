package handlers

import (
	"sg/services/public-api/controller"

	"github.com/gin-gonic/gin"
)

// TODO: Only has limited set of handler methods (which map)
type TestHandler interface {
	ListSilos(ctx *gin.Context)
	GetSilo(ctx *gin.Context)
}

// TODO: Can successfully limit handler methods but can't pass in a more limited controller
// 	Would like to pass in limited controller so compilation error will occur when the interface
//	attempts to use controller method that it does not have
func NewTestHandler(host string, conns int) TestHandler {
	cf := controller.NewControllerFactory(host, conns)
	return &APIHandler{
		Controller: cf.Controller(),
	}
}
