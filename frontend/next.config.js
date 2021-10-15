/** @type {import('next').NextConfig} */
module.exports = {
    reactStrictMode: true,
    images: {
        domains: [],
    },
    exportPathMap: async function () {
        // exporting only index for SPA redirectes 
        // (forcing index.html) (disabled-SSR)
        return {
          '/': { page: '/' }, 
        }
      },
}