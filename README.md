# slog-ecs

An [Elastic Common Schema (ECS)](https://www.elastic.co/guide/en/ecs/current/index.html) handler for Go's `log/slog` package, with an optional HTTP request logging middleware.

## Requirements

Go 1.21+

## Installation

```sh
go get github.com/unioslo/slog-ecs
```

## Usage

### Handler

```go
import (
    "log/slog"
    "github.com/unioslo/slog-ecs"
)

slog.SetDefault(slog.New(slogecs.NewHandler(slog.LevelInfo)))

slog.Info("application started")
```

Output:

```json
{
  "@timestamp": "2026-03-19T10:00:00.000Z",
  "log": { "level": "info" },
  "message": "application started",
  "ecs": { "version": "1.6.0" }
}
```

### HTTP middleware

```go
import (
    "net/http"
    "github.com/unioslo/slog-ecs"
)

handler := slogecs.HTTPLogger(mux)
```

Each request is logged as:

```json
{
  "@timestamp": "2026-03-19T10:00:00.000Z",
  "log": { "level": "info" },
  "message": "HTTP request",
  "http": {
    "request":  { "method": "GET" },
    "response": { "status_code": 200 }
  },
  "url":    { "path": "/api/example" },
  "event":  { "duration": 1234567 },
  "client": { "address": "10.0.0.1:54321" },
  "ecs":    { "version": "1.6.0" }
}
```

`event.duration` is in nanoseconds per the ECS specification.

## License

MIT — see [LICENSE](LICENSE).
