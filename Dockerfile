## THIS IS PRODUCTION DOCKERFILE.
## USED TO BUILD: slashbaseide/slashbase image

# Create base image for building go binary
FROM golang:1.17.3-alpine3.14 as base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=1

# System dependencies
RUN apk update && apk add --no-cache ca-certificates git build-base && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

# Executable builder
FROM base as backendbuilder

WORKDIR /slashbase
COPY . .
ENV GOOS="linux"
RUN go build --o backend -trimpath -ldflags="-X 'main.Build=production'"

# Install dependencies only when needed
FROM node:alpine AS deps

RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY ./frontend/package.json ./frontend/yarn.lock ./
RUN yarn install --frozen-lockfile

# Rebuild the source code only when needed
FROM node:alpine AS frontendbuilder

WORKDIR /app
COPY ./frontend/ .
COPY --from=deps /app/node_modules ./node_modules
RUN yarn build

# Production
FROM alpine:3.14

WORKDIR /slashbase
RUN mkdir -p /slashbase/data
COPY --from=backendbuilder /slashbase/backend /slashbase
COPY --from=frontendbuilder /app/out /slashbase/web

ENTRYPOINT ["/slashbase/backend"]
EXPOSE 3000