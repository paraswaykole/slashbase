## THIS IS FRONTEND DEVELOPMENT DOCKERFILE. 
## TO BE USED FOR LOCAL DEVELOPMENT.

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