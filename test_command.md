## metric module

### add prometheus module
```bash
curl -X POST localhost:31019/v1/metric-beat/modules -d '{"id":"prom", "module":"prometheus", "prometheus_module": {"period":10, "addr_list":["localhost:9000"]}'
```
