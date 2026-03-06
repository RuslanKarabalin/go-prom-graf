# backend

Base golang backend project setup.

## Libs

<https://github.com/spf13/viper> - for env/yaml config reading

```bash
go get github.com/spf13/viper
```

<https://github.com/jackc/pgx> - for PostgreSQL connection

```bash
go get github.com/jackc/pgx/v5
```

<https://github.com/redis/go-redis> - for Redis connection

```bash
go get github.com/redis/go-redis/v9
```

<https://github.com/open-telemetry/opentelemetry-go> - for logs, metrics, traces

```bash
go get go.opentelemetry.io/otel
```

<https://pkg.go.dev/go.opentelemetry.io/contrib/bridges/otelslog> - for opentelemetry and slog integration

```bash
go get go.opentelemetry.io/contrib/bridges/otelslog
```

## Tools

<https://github.com/oapi-codegen/oapi-codegen> - for generating code from openapi specification

<https://github.com/pressly/goose> - for database migration

<https://github.com/golangci/golangci-lint> - linter

## Project structure

```plaitext
backend/
├── api/                    # OpenAPI specs (.yaml), generated code goes here
├── bin/                    # local tool binaries (gitignored)
├── cmd/
│   └── api/
│       └── main.go         # entrypoint - wire up deps, start server
├── infra/                  # infrastructure config, not Go code
│   ├── docker-compose.yaml
│   ├── otelcol/
│   ├── prometheus/
│   └── grafana/
├── internal/               # private app code, not importable from outside
│   ├── config/             # config struct, viper loading
│   ├── handler/            # HTTP handlers (one file per domain)
│   ├── service/            # business logic
│   ├── repository/         # DB queries (pgx), one file per entity
│   ├── middleware/         # HTTP middleware (auth, logging, tracing)
│   └── model/              # domain types / DTOs
├── migrations/             # SQL migration files (goose)
├── Dockerfile
├── Makefile
└── go.mod
```

### Conventions

- `api/` - source-of-truth OpenAPI spec. Run `oapi-codegen` to generate server stubs and types into `internal/`.
- `bin/` - pinned tool binaries installed via `make install-*`. Checked into gitignore so everyone gets the same version via `go install`.
- `cmd/` - one subdirectory per binary. Keeps `main.go` thin: parse config, init deps, call into `internal/`.
- `infra/` - docker compose, OTel collector config, Prometheus, Grafana dashboards. Nothing here is compiled.
- `internal/` - everything the app owns. Go enforces that nothing outside this module can import it.
  - `handler/` calls `service/`, `service/` calls `repository/`. Never skip layers.
  - `repository/` only knows about the DB. No business logic here.
  - `model/` holds plain structs shared across layers. No methods with side effects.
- `migrations/` - goose SQL files named `00001_create_users.sql`, etc. Never edit applied migrations.
