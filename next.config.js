/** @type {import('next').NextConfig} */
module.exports = {
  reactStrictMode: true,
  env: {
        API_HOST: "http://localhost:3001",
        IS_LIVE: false,
    },
    images: {
        domains: [],
    },
}