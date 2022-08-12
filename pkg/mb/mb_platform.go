package mb

import (
	"github.com/paas-mate/gutil"
	"go.uber.org/zap"
	"metricbeat_mate/pkg/path"
	"metricbeat_mate/pkg/util"
)

func startMbPlatform() error {
	stdout, stderr, err := gutil.CallScript(path.MbStartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	return err
}

func restartMbPlatform() error {
	stdout, stderr, err := gutil.CallScript(path.MbRestartScript)
	util.Logger().Info("shell result ", zap.String("stdout", stdout), zap.String("stderr", stderr))
	return err
}
