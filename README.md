# tcp-lb

A Layer 4 TCP load balancer written in Go from scratch. Forwards raw TCP connections to a pool of backends without inspecting the payload, making it protocol-agnostic (HTTP, WebSocket, gRPC, etc.).

## Features

- Three load balancing strategies
- Per-backend health checking
- Graceful shutdown
- Docker and Docker Compose support

## Strategies

| Strategy | Description |
|---|---|
| `roundrobin` | Cycles through backends in order using an atomic counter |
| `leastconnection` | Routes to the backend with the fewest active connections |
| `iphash` | Hashes the client IP to a backend for sticky sessions, falls back to next healthy backend |

## Configuration

```yaml
listen_addr: ":8080"
backends:
  - ":9001"
  - ":9002"
  - ":9003"
strategy: "roundrobin"
```

## Running locally

```bash
go build -o tcplb .
./tcplb
```

Backends are expected to be running on the addresses defined in `config.yaml`.

## Running with Docker Compose

Production-like (proxy only, bring your own backends):
```bash
docker compose up --build
```

Local development (proxy + three socat echo backends for testing):
```bash
docker compose -f docker-compose.yaml -f docker-compose.dev.yaml up --build
```

Test the dev setup with:
```bash
echo "hello" | nc localhost 8080
```

## Running tests

```bash
go test ./...
```