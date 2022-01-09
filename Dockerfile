# Development with dlv Debugger
FROM golang:1.17.3-alpine3.14 as dev

WORKDIR /slashbase

COPY . .
RUN apk add --update alpine-sdk
RUN env GOOS=linux GOARCH=amd64 go build --o backend -trimpath
RUN go get github.com/go-delve/delve/cmd/dlv

ENTRYPOINT ["dlv"]
CMD ["debug","--headless", "--accept-multiclient", "--continue", "--log","-l" ,"0.0.0.0:2345" ,"--api-version=2", "--", "-e", "development"]
EXPOSE 3001
EXPOSE 2345

# Executable builder
FROM golang:1.17.3-alpine3.14 as builder

WORKDIR /slashbase
COPY . .
RUN apk add --update alpine-sdk
RUN env GOOS=linux GOARCH=amd64 go build --o backend -trimpath

# Production
FROM alpine:3.14

WORKDIR /slashbase
COPY --from=builder /slashbase/backend /slashbase

ENTRYPOINT ["/slashbase/backend"]
CMD ["-e", "production"]
EXPOSE 3001