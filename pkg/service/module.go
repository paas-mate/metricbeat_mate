package service

import (
	"bufio"
	"go.uber.org/zap"
	"metricbeat_mate/pkg/mb/generate"
	"metricbeat_mate/pkg/module"
	"metricbeat_mate/pkg/path"
	"metricbeat_mate/pkg/util"
	"os"
	"path/filepath"
)

func AddModule(req module.MetricModuleAddReq) error {
	file, err := os.OpenFile(getModuleFileName(req.Id), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(generate.ConvModule(req))
	if err != nil {
		util.Logger().Error("write string failed", zap.Error(err))
		return err
	}
	return writer.Flush()
}

func DelModule(moduleId string) error {
	return os.Remove(getModuleFileName(moduleId))
}

func getModuleFileName(moduleId string) string {
	return filepath.FromSlash(path.MbModulesDir + "/" + moduleId + ".yml")
}
