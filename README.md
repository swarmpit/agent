# Swarmpit agent

Swarmpit docker agent.

[![version](https://img.shields.io/github/release-pre/swarmpit/agent.svg)](https://github.com/swarmpit/agent/releases) 
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/swarmpit/agent/pulls)

## Run

```{r, engine='bash', count_lines}
docker run -d \
  --name agent \
  --volume /var/run/docker.sock:/var/run/docker.sock \
  swarmpit/agent:latest
```

### Parameters

- STATS_FREQUENCY - default to **30**
- EVENT_ENDPOINT - default to **http://app:8080/events**
- HEALTH_CHECK_ENDPOINT - default to **http://app:8080/version**
- DISK_PATH - mountpoint reported as node disk usage, default to **/** (**\\** on windows)
- DEBUG_EVENT - default to **false**
- DEBUG_STATS - default to **false**

The docker client is configured from the standard docker environment variables:

- DOCKER_HOST - docker endpoint, default to the mounted **/var/run/docker.sock**. Set to e.g. **tcp://proxy:2375** to go through a socket proxy instead of mounting the socket
- DOCKER_API_VERSION - docker api version used by the client. Must be supported by your engine, docker 29 dropped everything below **1.44**
- DOCKER_CERT_PATH, DOCKER_TLS_VERIFY - tls settings for a remote endpoint

## Important!

In case you are deploying agent inside Swarmpit [stack](https://github.com/swarmpit/swarmpit/blob/master/docker-compose.yml)
with some sort of customization, make sure that Swarmpit service name (default to **app**) match domain name set by EVENT_ENDPOINT & HEALTH_CHECK_ENDPOINT. 