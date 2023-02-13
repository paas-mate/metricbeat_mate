#!/bin/bash

mkdir -p $METRICBEAT_HOME/logs
CONF_FILE=$METRICBEAT_HOME/metricbeat.yml
echo 'hostfs: "/hostfs"' >$CONF_FILE
echo "metricbeat.config.modules:" >$CONF_FILE
echo "  path: ${path.config}/modules.d/*.yml" >>$CONF_FILE
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
if [ ${HOST_METADATA_ENABLED:-true} ]; then
   echo "  - add_host_metadata: ~" >>$CONF_FILE
fi
if [ ${CLOUD_METADATA_ENABLED:-true} ]; then
   echo "  - add_cloud_metadata: ~" >>$CONF_FILE
fi
if [ ${DOCKER_MATEDATA_ENABLED:-true} ]; then
   echo "  - add_docker_metadata: ~" >>$CONF_FILE
fi
if [ ${KUBERNETES_METADATA_ENABLED} ]; then
   echo "  - add_kubernetes_metadata: ~" >>$CONF_FILE
fi
nohup $METRICBEAT_HOME/metricbeat -e -system.hostfs=/hostfs -c $CONF_FILE 2>>$METRICBEAT_HOME/logs/metricbeat.log &
