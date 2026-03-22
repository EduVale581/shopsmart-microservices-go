FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY inventory-service ./inventory-service

RUN mkdir -p /app/bin && go build -o /app/bin/inventory-service ./inventory-service

EXPOSE 3002

CMD ["/app/bin/inventory-service"]
