ARG GO_VERSION=1.14.2
FROM golang:$GO_VERSION

# Caching packages.
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download

COPY cache.go .
RUN go build -tags cache cache.go

COPY . /src
RUN go build -o /usr/bin/api ./cmd

ENTRYPOINT ["api"]
