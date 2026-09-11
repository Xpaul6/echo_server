FROM golang:1.26.5 AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o echo_server

FROM alpine:3.24
WORKDIR /root/
COPY --from=builder /app/echo_server ./
EXPOSE 8080
CMD [ "./echo_server" ]
