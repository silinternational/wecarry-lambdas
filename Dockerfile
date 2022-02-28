FROM golang:latest

# Install packages
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get install -y git nodejs netcat

# Copy in source and install deps
COPY ./package.json /build
WORKDIR /
RUN npm install -g serverless && npm install
COPY ./ /src/
RUN go get ./...
