package path

import (
	"os"
	"path/filepath"
)

// metricbeat
var (
	MbHome       = os.Getenv("METRICBEAT_HOME")
	MbConfig     = filepath.FromSlash(MbHome + "/metricbeat.yml")
	MbModulesDir = filepath.FromSlash(MbHome + "/modules.d")
)

// mate
var (
	MbMatePath      = filepath.FromSlash(MbHome + "/mate")
	MbScripts       = filepath.FromSlash(MbMatePath + "/scripts")
	MbStartScript   = filepath.FromSlash(MbScripts + "/start-metricbeat.sh")
	MbRestartScript = filepath.FromSlash(MbScripts + "/restart-metricbeat.sh")
)
