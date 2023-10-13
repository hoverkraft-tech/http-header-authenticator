# http-header-authenticator

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
helm repo add hoverkraft-tech harbor.hoverkraft.cloud/public/charts
helm repo update hoverkraft-tech
helm upgrade http-header-authenticator hoverkraft-tech/http-header-authenticator \
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
