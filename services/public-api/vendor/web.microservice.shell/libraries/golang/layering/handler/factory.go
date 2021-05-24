package handler

import (
	"web.microservice.shell/libraries/golang/layering/query/psql_conn"
	"web.microservice.shell/libraries/golang/layering/query/sg_conn"
	"web.microservice.shell/libraries/golang/layering/controller"
)

func NewPublicAPIHandler(psqlArgs *psql_conn.StoreArgs, sgArgs *sg_conn.StoreArgs) PublicAPIHandler {
	cf := controller.NewControllerFactory(psqlArgs, sgArgs)
	return &BaseHandler{
		controller: cf.Controller(),
	}
}
