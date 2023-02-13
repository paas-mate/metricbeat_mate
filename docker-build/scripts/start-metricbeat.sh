#!/bin/bash

mkdir -p $METRICBEAT_HOME/logs
CONF_FILE=$METRICBEAT_HOME/metricbeat.yml
echo 'hostfs: "/hostfs"' >$CONF_FILE
echo "metricbeat.config.modules:" >$CONF_FILE
echo "- module: system" >>$CONF_FILE
echo "  period: 10s" >>$CONF_FILE
echo "  metricsets:" >>$CONF_FILE
echo "    - cpu" >>$CONF_FILE
echo "    - load" >>$CONF_FILE
echo "    - memory" >>$CONF_FILE
echo "    - network" >>$CONF_FILE
echo "    - process" >>$CONF_FILE
echo "    - process_summary" >>$CONF_FILE
echo "    - socket_summary" >>$CONF_FILE
echo "    - socket" >>$CONF_FILE
echo "  process.include_top_n:" >>$CONF_FILE
echo "    by_cpu: ${ES_CPU:-5}" >>$CONF_FILE
echo "    by_memory: ${ES_MEMORY:-5}" >>$CONF_FILE
echo "- module: system" >>$CONF_FILE
echo "  period: 1m" >>$CONF_FILE
echo "  metricsets:" >>$CONF_FILE
echo "    - filesystem" >>$CONF_FILE
echo "    - fsstat" >>$CONF_FILE
echo "  processors:" >>$CONF_FILE
echo "  - drop_event.when.regexp:" >>$CONF_FILE
echo "      system.filesystem.mount_point: '^/(sys|cgroup|proc|dev|etc|host|lib|snap) ($|/)'" >>$CONF_FILE
echo "- module: system" >>$CONF_FILE
echo "  period: 15m" >>$CONF_FILE
echo "  metricsets:" >>$CONF_FILE
echo "    - uptime" >>$CONF_FILE
echo "setup.template.settings:" >>$CONF_FILE
echo "  index.number_of_shards: 1" >>$CONF_FILE
echo "  index.codec: best_compression" >>$CONF_FILE
echo "setup.kibana:" >>$CONF_FILE
echo "  host: ${KIBANA_HOST:-locahost:5601}" >>$CONF_FILE
echo "output.elasticsearch:" >>$CONF_FILE

echo -n '  hosts: ["' >>$CONF_FILE
echo -n ${ES_HOST:-locahost:9200} >> $CONF_FILE
echo '"]' >>$CONF_FILE

echo "processors:" >>$CONF_FILE
echo "  - add_host_metadata: ~" >>$CONF_FILE
echo "  - add_cloud_metadata: ~" >>$CONF_FILE
echo "  - add_docker_metadata: ~" >>$CONF_FILE
echo "  - add_kubernetes_metadata: ~" >>$CONF_FILE
nohup $METRICBEAT_HOME/metricbeat -e -system.hostfs=/hostfs -c $CONF_FILE >>$METRICBEAT_HOME/logs/metricbeat.stdout.log 2>>$METRICBEAT_HOME/logs/metricbeat.stderr.log &
