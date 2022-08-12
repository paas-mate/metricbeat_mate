package config

import "github.com/paas-mate/gutil"

// es config
var (
	EsHost     string
	KibanaHost string
)

func init() {
	EsHost = gutil.GetEnvStr("ES_HOST", "localhost")
	KibanaHost = gutil.GetEnvStr("KIBANA_HOST", "localhost")
}
