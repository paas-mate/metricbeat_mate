FROM ttbb/metricbeat:nake

COPY docker-build /opt/metricbeat/mate

WORKDIR /opt/metricbeat

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/metricbeat/mate/scripts/start.sh"]
