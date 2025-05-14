
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

