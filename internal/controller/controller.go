package controller

import (
	"github.com/Chandan94f/mb-view/internal/dao"
	"github.com/Chandan94f/mb-view/internal/service"

	// "github.com/Chandan94f/mb-view/internal/utility/response"
	"github.com/Chandan94f/mb-utility/utility/response"
)

type Controller struct {
	resp    *response.ResponseHandler
	dao     dao.DaoService
	service service.Service
}

type ControllerParams struct {
	Resp    *response.ResponseHandler
	Dao     dao.DaoService
	Service service.Service
}

func NewController(controllerParams ControllerParams) *Controller {
	return &Controller{
		resp:    controllerParams.Resp,
		dao:     controllerParams.Dao,
		service: controllerParams.Service,
	}
}
