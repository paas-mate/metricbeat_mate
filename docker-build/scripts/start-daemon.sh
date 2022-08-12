#!/bin/bash

mkdir -p $METRICBEAT_HOME/logs
nohup $METRICBEAT_HOME/mate/metricbeat_mate >>$METRICBEAT_HOME/logs/metricbeat_mate.stdout.log 2>>$METRICBEAT_HOME/logs/metricbeat_mate.stderr.log
