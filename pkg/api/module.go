package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metricbeat_mate/pkg/module"
	"metricbeat_mate/pkg/service"
	"metricbeat_mate/pkg/util"
	"net/http"
)

func AddModule(c *gin.Context) {
	req := module.MetricModuleAddReq{}
	err := c.BindJSON(&req)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	err = service.AddModule(req)
	if err != nil {
		util.Logger().Error("route add error ", zap.Error(err))
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusCreated)
}

func DelModule(c *gin.Context) {
	id := c.Param("id")
	err := service.DelModule(id)
	if err != nil {
		util.Logger().Error("route delete error ", zap.Error(err))
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusNoContent)
}
