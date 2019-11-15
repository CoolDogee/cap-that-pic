FROM golang:alpine AS builder
ADD . /app
WORKDIR /app/go/
RUN go build -o main .

FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install
RUN npm run build

FROM alpine:latest
COPY --from=builder /app/go/main ./
COPY --from=node_builder /build ./web
COPY --from=builder /app/lyrics/lyrics.json ./lyrics.json
RUN chmod +x ./main
EXPOSE 8080

CMD ./main