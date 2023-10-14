ci: go-tests go-cover golang-ci

go-tests: src/* go.mod go.sum
	go test -v ./...

go-cover: src/* go.mod go.sum
	go test -coverprofile=coverage.out ./...

golang-ci: src/* go.mod go.sum
	docker run --rm -t \
		-v ${PWD}:/app \
		-v ${PWD}/.cache/golang-ci:/root/.cache \
		-w /app golangci/golangci-lint:v1.54.2 golangci-lint run
