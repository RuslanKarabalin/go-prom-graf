FROM golang:alpine3.23 AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN go mod tidy
RUN go build -o myapp main.go

FROM golang:alpine3.23
WORKDIR /app
COPY --from=builder /app/myapp myapp
EXPOSE 2112
ENTRYPOINT ["./myapp"]
