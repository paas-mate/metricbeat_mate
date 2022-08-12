FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o metricbeat_mate .


FROM ttbb/metricbeat:nake

COPY docker-build /opt/sh/metricbeat/mate

COPY --from=build /opt/sh/compile/pkg/metricbeat_mate /opt/sh/metricbeat/mate/metricbeat_mate

WORKDIR /opt/sh/metricbeat

CMD ["/usr/bin/dumb-init", "bash", "-vx", "/opt/sh/metricbeat/mate/scripts/start.sh"]
