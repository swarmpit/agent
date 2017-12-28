# event-collector

Docker event collector.

## Run

```{r, engine='bash', count_lines}
docker run -d \
  --name ec \
  --env EVENT_ENDPOINT=http://endpoint_to_send_events \
  --volume /var/run/docker.sock:/var/run/docker.sock \
  swarmpit/event-collector:latest
```
