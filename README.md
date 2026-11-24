# http-header-authenticator

<div align="center">
  <img src="https://github.com/hoverkraft-tech.png" alt="Hoverkraft logo" width="200" />
</div>

---

[![Main - Continuous Integration](https://github.com/hoverkraft-tech/http-header-authenticator/actions/workflows/main-ci.yml/badge.svg)](https://github.com/hoverkraft-tech/http-header-authenticator/actions/workflows/main-ci.yml)
[![GitHub release](https://img.shields.io/github/v/release/hoverkraft-tech/http-header-authenticator)](https://github.com/hoverkraft-tech/http-header-authenticator/releases)
[![PRs welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

## Overview

Go service that authenticates HTTP requests against a configured header and value.

## Usage

```sh
http-header-authenticator check -H 'X-acme-auth' -V 'SomeSecurePassword'
```

```sh
curl -H 'X-acme-auth: SomeSecurePassword' localhost:8080/check # -> 200
curl localhost:8080/check # -> 403
```

Use it as an external authentication service for an ingress controller, for example
to authenticate requests coming from a CDN. `/health` reports service readiness.

## Installation

### Helm

```sh
helm upgrade http-header-authenticator oci://harbor.hoverkraft.cloud/public/charts/http-header-authenticator \
  --install --create-namespace --namespace http-header-authenticator \
  --set-string securityHeader.name=X-Auth-Token \
  --set-string securityHeader.value="$AUTH_TOKEN"
```

Set `AUTH_TOKEN` to your authentication secret before installing. The chart runs as
a non-root user with a read-only filesystem and resource budgets. Its NetworkPolicy
allows incoming HTTP traffic and denies outbound connections. Adjust these values
for your cluster and workload.

### Docker

- [Docker Hub](https://hub.docker.com/repository/docker/webofmars/http-header-authenticator)

```shell
AUTH_TOKEN='replace-with-a-strong-secret' docker compose up -d --build
```

### From source

```shell
go build -o ./bin/http-header-authenticator ./src/github.com/hoverkraft-tech/http-header-authenticator
```

See the [development guide](docs/development/README.md) for local checks.

## Contributing

Contributions are welcome. Read [Contributing](CONTRIBUTING.md) before proposing a change.
