FROM golang:1.22.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags "-s -w" -o server ./cmd/main.go


FROM gcr.io/distroless/base-debian12 AS runner
WORKDIR /app

COPY --from=builder /app/server ./

ENV GO_ENV=production

EXPOSE 3000

CMD ["./server"]