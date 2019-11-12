FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app/go/
RUN go build -o main .

FROM node:alpine AS node_builder
WORKDIR /app/client
RUN npm install
RUN npm run build
COPY /app/client/build /app/

CMD ./main