## THIS IS PRODUCTION DOCKERFILE.
## USED TO BUILD: slashbaseide/slashbase image

# Create base image for building go binary
FROM golang:1.20.3-alpine3.17 as base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=1

# System dependencies
RUN apk update && apk add --no-cache ca-certificates git build-base && update-ca-certificates

COPY go.mod go.sum ./
RUN go mod download

# Install dependencies only when needed
FROM node:alpine AS deps

RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY ./frontend/server/package.json ./frontend/server/yarn.lock ./
RUN yarn install --frozen-lockfile

# Rebuild the source code only when needed
FROM node:alpine AS frontendbuilder

WORKDIR /app
COPY ./frontend/server/ .
COPY --from=deps /app/node_modules ./node_modules
RUN yarn build-server

# Executable builder
FROM base as backendbuilder

WORKDIR /app
COPY . .
COPY --from=frontendbuilder /dist /app/frontend/dist
RUN make build-server

# Production
FROM alpine:3.14

WORKDIR /slashbase
COPY --from=backendbuilder /app/slashbase /slashbase

ENTRYPOINT ["/slashbase/slashbase"]
EXPOSE 3000