# Stage 1: Build the Go app
FROM golang:1.24.6-alpine AS build

WORKDIR /app

ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY internal/config/config.yaml config/

RUN go build -o /elearning ./cmd/server

# Stage 2: Run the Go app
FROM alpine:latest

RUN apk --no-cache add ca-certificates bash

WORKDIR /app

COPY --from=build /elearning /elearning

# Add wait-for-it.sh and start.sh
COPY wait-for-it.sh /wait-for-it.sh
COPY start.sh /start.sh
RUN chmod +x /elearning /wait-for-it.sh /start.sh

EXPOSE 8080

# Run the start.sh script
CMD ["/start.sh"]

