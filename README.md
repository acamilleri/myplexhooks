# MyPlexHooks

This is my personal Plex Hooks repository

**Plex Webhooks require a Plex Pass subscription**

You can find more details in [examples](examples) folder

## Installation

## Usage

requirements
```
export PLEXHOOKS_GRAFANA_URL=https://mygrafana.fr
export PLEXHOOKS_GRAFANA_TOKEN=secret
```

### Simple
```bash
go run --log.formatter json
```

### Docker

```bash
docker run -it --env PLEXHOOKS_GRAFANA_URL --env PLEXHOOKS_GRAFANA_TOKEN acamilleri/myplexhooks:latest --log.formatter json
```
