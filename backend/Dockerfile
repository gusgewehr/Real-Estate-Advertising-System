FROM golang:1.25-alpine AS builder

RUN apk add --no-cache tzdata

RUN apk add --no-cache ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/api/main.go

RUN go build -o main ./cmd/api/main.go

FROM scratch


COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]