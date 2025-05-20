
run:
	@go run .

build:
	@echo "> Building binary"
	@go build

clean:
	@go clean

lint:
	@echo "  > Linting code"
	@golangci-lint run

fmt:
	@echo "  > Formatting code"
	@go fmt ./...

test:
	@echo "  > Executing unit tests"
	@go test -v -timeout 60s -race ./...

coverage:
	@echo "  > Running test test and coverage"
	test
	go tool cover -html=coverage.out

vet:
	@echo "  > Checking code with vet"
	@go vet ./...

tidy:
	@echo "  > Running tidy to fix dependencies"
	@go mod tidy


setup:
	@echo "  > Download dependencies..."
	@go mod download && go mod tidy
	@echo "  > Dependencies downloaded."

init:
	 setup

# Load testing
server:
	@echo "  > Starting API server on port 8080"
	@go run . --server --port=:8080

load-test:
	@echo "  > Running K6 load test"
	@k6 run load_test.js

# Docker commands
docker-build:
	@echo "  > Building Docker image. See Dockerfile"
	@docker build -t generate-person-data-json .

docker-run-server:
	@echo "  > Running Docker container in server mode"
	@docker run -p 8080:8080 generate-person-data-json

docker-run-generate:
	@echo "  > Running Docker container to generate data"
	@docker run -v $$(pwd):/output generate-person-data-json --server=false -o /output/fake_pii_data.json
