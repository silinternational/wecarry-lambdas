FROM golang:latest

# Install packages
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get install -y git nodejs netcat

# Copy in source and install deps
RUN mkdir -p /build
COPY ./package.json /build
WORKDIR /build
RUN npm install -g serverless && npm install
COPY ./ /build
RUN go get ./...
