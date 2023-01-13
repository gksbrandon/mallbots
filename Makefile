# Installing Tools and Generating
install-tools:
	echo installing tools && \
	go install \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	echo done

generate:
	echo running code generation
	go generate ./...
	echo done

# Docker Compose
dc-up:
	docker-compose up -d 

dc-down:
	docker-compose down --rmi all

dc-reset: dc-down dc-up

dc-update:
	docker-compose down
	docker-compose build monolith
	docker-compose up -d

dc-logs:
	docker-compose logs -f
