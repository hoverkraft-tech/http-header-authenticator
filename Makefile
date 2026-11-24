ci: go-tests go-cover golang-ci

go-tests: src/* go.mod go.sum
	go test -race -count=1 -timeout=5m ./...

go-cover: src/* go.mod go.sum
	go test -race -count=1 -timeout=5m -covermode=atomic -coverprofile=coverage.out ./...

golang-ci: src/* go.mod go.sum
	docker run --rm -t \
		-v ${PWD}:/app \
		-v ${PWD}/.cache/golang-ci:/root/.cache \
		-w /app golangci/golangci-lint:v2.15.0 golangci-lint run
