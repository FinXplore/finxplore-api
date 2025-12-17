# Variables
APP_NAME=finxplore-api
CMD_PATH=./cmd/server

.PHONY: run build clean docker-up docker-down

# Run locally (wire)
wire:
	wire $(CMD_PATH)
run:
	go run $(CMD_PATH)

# Build binary
build:
	
	go build -o bin/$(APP_NAME) $(CMD_PATH)

# Clean build artifacts
clean:
	rm -rf bin/

# Docker: Start everything
up:
	docker builder prune -a -f
	docker-compose build --no-cache api
	docker-compose up --build

# Docker: Stop everything
down:
	docker-compose down

# Docker: View logs
logs:
	docker-compose logs -f

# Install gRPC tools
init-grpc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Go code from Proto files
proto:
	protoc --proto_path=api/proto/v1 \
	--go_out=pkg/pb/v1 --go_opt=paths=source_relative \
	--go-grpc_out=pkg/pb/v1 --go-grpc_opt=paths=source_relative \
	api/proto/v1/*.proto