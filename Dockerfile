FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app/go/
RUN go build -o main .
COPY --from=builder /main ./

FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install
RUN npm run build

CMD ./main