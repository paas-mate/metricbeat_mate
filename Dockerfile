FROM ttbb/base:go AS build
COPY . /opt/compile
WORKDIR /opt/compile/pkg
RUN go build -o metricbeat_mate .


FROM ttbb/metricbeat:nake

COPY docker-build /opt/metricbeat/mate

COPY --from=build /opt/compile/pkg/metricbeat_mate /opt/metricbeat/mate/metricbeat_mate

WORKDIR /opt/metricbeat

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/metricbeat/mate/scripts/start.sh"]
