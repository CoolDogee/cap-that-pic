# Build the Go API
FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app/go/
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/go/main"]
CMD ./main