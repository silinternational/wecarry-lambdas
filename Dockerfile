FROM golang:1.18

RUN curl -o- -L https://slss.io/install | VERSION=3.8.0 bash

# Copy in source and install deps
WORKDIR /src
COPY ./ /src/
RUN go get ./...
