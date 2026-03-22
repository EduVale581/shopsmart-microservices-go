FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY orders-service ./orders-service

RUN mkdir -p /app/bin && go build -o /app/bin/orders-service ./orders-service

EXPOSE 3003

CMD ["/app/bin/orders-service"]
