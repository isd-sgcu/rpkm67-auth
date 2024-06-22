server:
	go run cmd/main.go

watch: 
	air

mock-gen:
	mockgen -source ./internal/auth/auth.repository.go -destination ./mocks/auth/auth.repository.go
	mockgen -source ./internal/auth/auth.service.go -destination ./mocks/auth/auth.service.go
	mockgen -source ./internal/user/user.service.go -destination ./mocks/user/user.service.go

test:
	go vet ./...
	go test  -v -coverpkg ./internal/... -coverprofile coverage.out -covermode count ./internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

proto:
	go get github.com/isd-sgcu/johnjud-go-proto@latest