FROM golang:1.17

RUN curl -o- -L https://slss.io/install | VERSION=3.3.0 bash && \
    mv $HOME/.serverless/bin/serverless /usr/local/bin && \
    ln -s /usr/local/bin/serverless /usr/local/bin/sls

# Copy in source and install deps
WORKDIR /src
COPY ./ /src/
RUN go get ./...
