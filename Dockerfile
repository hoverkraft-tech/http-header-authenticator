# build stage
FROM golang:1.27-alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY go.* .
COPY src/ .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download && go mod verify
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -mod=readonly -trimpath -o /go/bin/http-header-authenticator -v ./...

# final stage
FROM alpine:3.18
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/http-header-authenticator /usr/local/bin/http-header-authenticator
COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]
LABEL "name"="http-header-authenticator" "version"="0.1.0"
EXPOSE 8080
ENV GIN_MODE=release
