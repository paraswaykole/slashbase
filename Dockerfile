## THIS IS BACKEND DEVELOPMENT DOCKERFILE. 
## TO BE USED FOR LOCAL DEVELOPMENT.

# Create base image for building go binary
FROM golang:1.17.3-alpine3.14 as base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

FROM base as debugger
WORKDIR /slashbase
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Development with dlv Debugger
FROM debugger as dev

WORKDIR /slashbase

ENTRYPOINT ["dlv"]
CMD ["debug","--headless", "--accept-multiclient", "--continue", "--log","-l" ,"0.0.0.0:2345" ,"--api-version=2", "--", "-e", "development"]
EXPOSE 3001
EXPOSE 2345