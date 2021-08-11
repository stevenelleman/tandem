package psql_sg_handler

import (
	"web.microservice.shell/libraries/golang/layering/controller/psql_sg_controller"
	"web.microservice.shell/libraries/golang/layering/query/psql_conn"
	"web.microservice.shell/libraries/golang/layering/query/sg_conn"
)

func NewHandler(psqlArgs *psql_conn.StoreArgs, sgArgs *sg_conn.StoreArgs) *Handler {
	cf := psql_sg_controller.NewControllerFactory(psqlArgs, sgArgs)
	return &Handler{
		controller: cf.Controller(),
	}
}
