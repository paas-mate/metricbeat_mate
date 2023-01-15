package generate

import (
	"encoding/json"
	"metricbeat_mate/pkg/config"
	"metricbeat_mate/pkg/module"
	"strconv"
	"strings"
)

func ConvMb() string {
	var sb strings.Builder
	sb.WriteString("metricbeat.config.modules:\n")
	sb.WriteString("  path: ${path.config}/modules.d/*.yml\n")
	sb.WriteString("  reload.enabled: true\n")
	sb.WriteString("setup.template.settings:\n")
	sb.WriteString("  index.number_of_shards: 1\n")
	sb.WriteString("  index.codec: best_compression\n")
	sb.WriteString("setup.kibana:\n")
	sb.WriteString(ConvKibanaAddr())
	sb.WriteString("output.elasticsearch:\n")
	sb.WriteString(ConvEsAddr())
	sb.WriteString("processors:\n")
	sb.WriteString("  - add_host_metadata: ~\n")
	sb.WriteString("  - add_cloud_metadata: ~\n")
	sb.WriteString("  - add_docker_metadata: ~\n")
	sb.WriteString("  - add_kubernetes_metadata: ~\n")
	return sb.String()
}

func ConvModule(req module.MetricModuleAddReq) string {
	var sb strings.Builder
	sb.WriteString("- module: ")
	sb.WriteString(req.Module)
	sb.WriteString("\n")
	if req.Module == "prometheus" {
		sb.WriteString(ConvPromModule(req.PrometheusModule))
	}
	return sb.String()
}

func ConvPromModule(module module.PrometheusModule) string {
	var sb strings.Builder
	// period
	sb.WriteString("  period: ")
	sb.WriteString(strconv.Itoa(module.Period))
	sb.WriteString("s")
	sb.WriteString("\n")
	sb.WriteString("  hosts: ")
	bytes, err := json.Marshal(module.AddrList)
	if err != nil {
		panic(err)
	}
	sb.WriteString(string(bytes))
	sb.WriteString("\n")
	if module.IncludeMetrics != nil {
		sb.WriteString("  metrics_filters:")
		sb.WriteString("\n")
		sb.WriteString("    include: [")
		sb.WriteString(strings.Join(module.IncludeMetrics, ","))
		sb.WriteString("]")
		sb.WriteString("\n")
	}
	return sb.String()
}

func ConvEsAddr() string {
	return "  hosts: [\"" + config.EsHost + "\"]\n"
}

func ConvKibanaAddr() string {
	return "  host: " + config.KibanaHost
}
