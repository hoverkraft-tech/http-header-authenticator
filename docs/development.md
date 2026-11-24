# Development

Use the Go toolchain selected by `go.mod` and golangci-lint v2.15.0.

```sh
go mod download
go mod verify
go mod tidy -diff
golangci-lint run
go test -race -count=1 -timeout=5m -covermode=atomic -coverprofile=coverage.out ./...
go run golang.org/x/vuln/cmd/govulncheck@v1.10.0 ./...
```

CI uses the setup, lint, test and audit actions from `ci-github-go` at `main`
while they are being validated in consumer repositories. The test action
produces JUnit and coverage reports. The same coverage profile is sent to
Coveralls without rerunning tests or converting formats.

Image builds wait for these checks. The existing container authentication and
Helm checks validate the built image before it can be promoted for release.
