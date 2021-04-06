package handlers

import (
	"web-microservice-shell/libraries/golang/guts/connection/service/psql_conn"
	"web-microservice-shell/libraries/golang/guts/connection/service/sg_conn"

	"web-microservice-shell/libraries/golang/guts/controller"
)

func NewPublicAPIHandler(psqlArgs *psql_conn.StoreArgs, sgArgs *sg_conn.StoreArgs) PublicAPIHandler {
	cf := controller.NewControllerFactory(psqlArgs, sgArgs)
	return &BaseHandler{
		controller: cf.Controller(),
	}
}
