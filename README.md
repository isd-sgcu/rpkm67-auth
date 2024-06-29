# rpkm67-auth

## Stack

-   golang
-   gRPC
-   postgresql
-   redis
-   minio

## Getting Started

### Prerequisites

-   💻
-   golang 1.22 or [later](https://go.dev)
-   docker
-   makefile
-   [Go Air](https://github.com/air-verse/air)

### Installation

1. Clone this repo
2. Run `go mod download` to download all the dependencies.

### Running only this service
1. Copy `.env.template` and paste it in the same directory as `.env`. Fill in the appropriate values.
2. Run `make docker`.
3. Run `make server` or `air` for hot-reload.

### Running all RPKM67 services (all other services are run as containers)
1. Copy `docker-compose.qa.template.yml` and paste it in the same directory as `docker-compose.qa.yml`. Fill in the appropriate values.
2. Run `make pull-latest-mac` or `make pull-latest-windows` to pull the latest images of other services.
1. Run `make docker-qa`.
2. Run `make server` or `air` for hot-reload.

### Unit Testing
1. Run `make test`

## Other microservices/repositories of RPKM67
- [gateway](https://github.com/isd-sgcu/rpkm67-gateway): Routing and request handling
- [auth](https://github.com/isd-sgcu/rpkm67-auth): Authentication and user service
- [backend](https://github.com/isd-sgcu/rpkm67-backend): Group, Baan selection and Stamp, Pin business logic
- [checkin](https://github.com/isd-sgcu/rpkm67-checkin): Checkin for events service
- [store](https://github.com/isd-sgcu/rpkm67-store): Object storage service for user profile pictures
- [model](https://github.com/isd-sgcu/rpkm67-model): SQL table schema and models
- [proto](https://github.com/isd-sgcu/rpkm67-proto): Protobuf files generator
- [go-proto](https://github.com/isd-sgcu/rpkm67-go-proto): Generated protobuf files for golang
- [frontend](https://github.com/isd-sgcu/firstdate-rpkm67-frontend): Frontend web application
