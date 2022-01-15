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

# Executable builder
FROM base as builder

WORKDIR /slashbase
COPY . .
RUN env GOOS=linux GOARCH=amd64 go build --o backend -trimpath

# Production
FROM alpine:3.14

WORKDIR /slashbase
COPY --from=builder /slashbase/backend /slashbase

ENTRYPOINT ["/slashbase/backend"]
CMD ["-e", "production"]
EXPOSE 3001