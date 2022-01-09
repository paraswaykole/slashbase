# Install dependencies only when needed
FROM node:alpine AS deps

RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile

# Development
FROM node:alpine AS dev

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules

ENTRYPOINT ["yarn", "dev"] 
EXPOSE 3000

# Rebuild the source code only when needed
FROM node:alpine AS builder

WORKDIR /app
COPY . .
COPY --from=deps /app/node_modules ./node_modules
RUN yarn build

# Production
FROM nginx:1.21.4-alpine

RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /app/out /usr/share/nginx/html

ENTRYPOINT ["nginx", "-g", "daemon off;"]
EXPOSE 80
EXPOSE 443