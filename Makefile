pull-latest-mac:
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-gateway:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-auth:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-backend:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-checkin:latest
	docker pull --platform linux/x86_64 ghcr.io/isd-sgcu/rpkm67-store:latest

pull-latest-windows:
	docker pull ghcr.io/isd-sgcu/rpkm67-gateway:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-auth:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-backend:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-checkin:latest
	docker pull ghcr.io/isd-sgcu/rpkm67-store:latest

docker:
	docker-compose up

docker-qa:
	docker-compose -f docker-compose.qa.yml up

server:
	go run cmd/main.go

watch: 
	air

mock-gen:
	mockgen -source ./internal/user/user.service.go -destination ./mocks/user/user.service.go
	mockgen -source ./internal/user/user.repository.go -destination ./mocks/user/user.repository.go

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

proto:
	go get github.com/isd-sgcu/rpkm67-go-proto@latest

model:
	go get github.com/isd-sgcu/rpkm67-model@latest