package module

type MetricModuleAddReq struct {
	Id               string           `json:"id"`
	Module           string           `json:"module"`
	PrometheusModule PrometheusModule `json:"prometheus_module"`
}

type PrometheusModule struct {
	Period         int      `json:"period"`
	AddrList       []string `json:"addr_list"`
	IncludeMetrics []string `json:"include_metrics"`
}
