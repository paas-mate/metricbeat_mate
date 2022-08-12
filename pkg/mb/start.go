package mb

import (
	"bufio"
	"go.uber.org/zap"
	"metricbeat_mate/pkg/mb/generate"
	"metricbeat_mate/pkg/path"
	"metricbeat_mate/pkg/util"
	"os"
)

func Start() {
	startMb()
}

func startMb() {
	err := generateMbFile()
	if err != nil {
		util.Logger().Error("generate metricbeat config file failed ", zap.Error(err))
		return
	}
	err = startMbPlatform()
	if err != nil {
		util.Logger().Error("run start ngx scripts failed ", zap.Error(err))
		return
	}
}

func restartMb() {
	err := generateMbFile()
	if err != nil {
		util.Logger().Error("generate metricbeat config file failed ", zap.Error(err))
		return
	}
	err = restartMbPlatform()
	if err != nil {
		util.Logger().Error("run start ngx scripts failed ", zap.Error(err))
		return
	}
}

func generateMbFile() (err error) {
	file, err := os.OpenFile(path.MbConfig, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(generate.ConvMb())
	if err != nil {
		util.Logger().Error("write string failed", zap.Error(err))
		return
	}
	err = writer.Flush()
	return
}
