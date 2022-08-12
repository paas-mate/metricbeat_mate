#!/bin/bash

mkdir -p $METRICBEAT_HOME/logs
nohup $METRICBEAT_HOME/metricbeat -e -system.hostfs=/hostfs -c $METRICBEAT_HOME/metricbeat.yml >>$METRICBEAT_HOME/logs/metricbeat.stdout.log 2>>$METRICBEAT_HOME/logs/metricbeat.stderr.log &
