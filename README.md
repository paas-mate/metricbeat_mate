# metricbeat_mate

environment list

- `ES_HOST` elasticsearch address, default `localhost:9200`
- `KIBANA_HOST` kibana address, default `localhost:5601`

## run on docker
```bash
docker run --rm \
  --name=metricbeat \
  --user=root \
  --volume="/var/run/docker.sock:/var/run/docker.sock:ro" \
  --volume="/sys/fs/cgroup:/hostfs/sys/fs/cgroup:ro" \
  --volume="/proc:/hostfs/proc:ro" \
  --volume="/:/hostfs:ro" \
  ttbb/metricbeat:mate
```
