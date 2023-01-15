package config

import "github.com/paas-mate/gutil"

// es config
var (
	EsHost     string
	KibanaHost string
)

func init() {
	EsHost = gutil.GetEnvStr("ES_HOST", "localhost:9200")
	KibanaHost = gutil.GetEnvStr("KIBANA_HOST", "localhost:5601")
}
