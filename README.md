# agent

Swarmpit docker agent.

## Run

```{r, engine='bash', count_lines}
docker run -d \
  --name agent \
  --publish 8080:8080 \
  --env EVENT_ENDPOINT=http://endpoint_to_send_events \
  --env HEALTH_CHECK_ENDPOINT=http://endpoint_to_check_whether_receiver_up \
  --volume /var/run/docker.sock:/var/run/docker.sock \
  swarmpit/agent:latest
```
