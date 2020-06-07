package handlers

import (
	"github.com/gin-gonic/gin"

	"sg/services/public-api/controller"
	"sg/services/public-api/handlers/requests"
)

type SiloRequest struct {
	State string `json:"state"`
	// separate struct from silos bc requests need auth tokens eventually
}

func ListSilos(c *gin.Context) {
	silos, err := controller.ListSilos()
	if err != nil {
		ReturnError(c, 400, err)
		// TODO: Someone is going to forget to put the return statement - how to refactor to avoid?
		return
	}
	ReturnJSONList(c, 200, silos)
	return
}

func GetSilo(c *gin.Context) {
	id := c.Param("silo_id")
	silo, err := controller.GetSilo(id)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}
	ReturnJSON(c, 200, silo)
}

func CreateSilo(c *gin.Context) {
	// TODO: should the id be in the request rather than the path?
	var json SiloRequest
	err := c.BindJSON(&json)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}

	id := c.Param("silo_id")
	siloReq := &requests.Silo{
		Id: id,
		State: json.State,
	}

	err = controller.CreateSilo(siloReq)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}
	c.Status(200)
}

func UpdateSilo(c *gin.Context) {
	var json SiloRequest

	err := c.BindJSON(&json)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}

	id := c.Param("silo_id")
	siloReq := &requests.Silo{
		Id: id,
		State: json.State,
	}

	err = controller.UpdateSilo(siloReq)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}
	c.Status(200)
}

func DeleteSilo(c *gin.Context) {
	id := c.Param("silo_id")
	err := controller.DeleteSilo(id)
	if err != nil {
		ReturnError(c, 400, err)
		return
	}
	c.Status(200)
}
