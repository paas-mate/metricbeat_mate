#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
bash -x $DIR/start-metricbeat.sh
tail -f /dev/null
