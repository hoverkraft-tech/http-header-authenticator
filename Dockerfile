FROM golang:1.27-alpine@sha256:70bd7ddea8b6cbe5d6943cde279d6b67feb10cc924fb9100d1d8757c1575b280 AS builder
WORKDIR /go/src/app
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && go mod verify
COPY src/ .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -mod=readonly -trimpath -o /go/bin/http-header-authenticator ./...

FROM alpine:3.24@sha256:f59abe656b692d901e184be3445e084786602108599d2a51d229c0f5a48c71bf
RUN addgroup -g 10001 app && adduser -D -u 10001 -G app app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/http-header-authenticator /usr/local/bin/http-header-authenticator
COPY --chmod=755 entrypoint.sh /entrypoint.sh
USER 10001:10001
EXPOSE 8080
ENV GIN_MODE=release
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 CMD ["wget", "-q", "--spider", "http://127.0.0.1:8080/health"]
ENTRYPOINT ["/entrypoint.sh"]
