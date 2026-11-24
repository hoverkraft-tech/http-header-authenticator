# http-header-authenticator

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

See [Contributing](CONTRIBUTING.md) to develop or propose changes.
