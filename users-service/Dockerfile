FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY users-service ./users-service

RUN mkdir -p /app/bin && go build -o /app/bin/users-service ./users-service

EXPOSE 3001

CMD ["/app/bin/users-service"]
