ci: golang-ci

golang-ci: src/* go.mod go.sum
	docker run --rm -t \
		-v ${PWD}:/app \
		-v ${PWD}/.cache/golang-ci:/root/.cache \
		-w /app golangci/golangci-lint:v1.54.2 golangci-lint run
