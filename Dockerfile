# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint
# compared to Ubuntu
FROM golang:1.20-alpine

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

WORKDIR /app/pkg

COPY pkg ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /rest-api-server

EXPOSE 8090
CMD [/rest-api-server]
