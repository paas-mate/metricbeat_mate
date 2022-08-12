package main

import (
	"github.com/gin-gonic/gin"
	"metricbeat_mate/pkg/api"
	"metricbeat_mate/pkg/mb"
	"metricbeat_mate/pkg/util"
)

func main() {
	util.Logger().Debug("this is a debug msg")
	util.Logger().Info("this is a info msg")
	util.Logger().Error("this is a error msg")
	mb.Start()
	router := gin.Default()
	router.POST("/v1/metric-beat/modules", api.AddModule)
	router.DELETE("/v1/metric-beat/modules/:id", api.DelModule)
	util.Logger().Info("metricbeat mate started")
	err := router.Run(":31019")
	if err != nil {
		util.Logger().Error("open http server failed")
		return
	}
}
