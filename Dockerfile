FROM golang:1.24-alpine

WORKDIR /app
COPY . .

RUN go build -o queue-service main.go

EXPOSE 8080
CMD ["./queue-service"]