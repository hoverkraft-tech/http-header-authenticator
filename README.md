# http-header-authenticator

![GitHub Verified Creator](https://img.shields.io/badge/GitHub-Verified%20Creator-4493F8?logo=github)
[![Continuous Integration](https://github.com/hoverkraft-tech/http-header-authenticator/actions/workflows/main-ci.yml/badge.svg)](https://github.com/hoverkraft-tech/http-header-authenticator/actions/workflows/main-ci.yml)
[![PRs welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)](https://github.com/hoverkraft-tech/http-header-authenticator/issues)
[![Release](https://img.shields.io/github/v/release/hoverkraft-tech/http-header-authenticator)](https://github.com/hoverkraft-tech/http-header-authenticator/releases)
[![Downloads](https://img.shields.io/github/downloads/hoverkraft-tech/http-header-authenticator/total)](https://github.com/hoverkraft-tech/http-header-authenticator/releases)
[![GitHub Stars](https://img.shields.io/github/stars/hoverkraft-tech/http-header-authenticator?style=social)](https://github.com/hoverkraft-tech/http-header-authenticator)

Go micro-service used to authenticate http requests based on provided fix http headers

## usage

```sh
http-header-authenticator check -H 'X-acme-auth' -V 'SomeSecurePassword'
```

```sh
curl -H 'X-acme-auth: SomeSecurePassword' localhost:8080 # -> 200
curl localhost:8080 # -> 404
```

Can be used with nginx-ingress controller to authenticate requests comming from CDN servers
See https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/configmap/#global-auth-url for inspiration

## installation

### using helm

```sh
helm upgrade http-header-authenticator oci://harbor.hoverkraft.cloud/public/charts/http-header-authenticator \
  --install --create-namespace --namespace http-header-authenticator
```

### docker

- [Docker Hub](https://hub.docker.com/repository/docker/webofmars/http-header-authenticator)

```shell
docker-compose up -d --build
```

### from sources

```shell
go build -o ./bin/http-header-authenticator ./src/github.com/hoverkraft-tech/http-header-authenticator
```
