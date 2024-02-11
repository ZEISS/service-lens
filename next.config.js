const path = require('path')

/** @type {import('next').NextConfig} */
const nextConfig = {
  output: 'standalone',
  trailingSlash: true,
  reactStrictMode: true,
  async rewrites() {
    return {
      fallback: [
        {
          source: '/:path*',
          destination: '/index.html'
        }
      ]
    }
  },
  experimental: {
    serverActions: {
      bodySizeLimit: '2mb'
    },
    serverComponentsExternalPackages: [
      '@trpc/server',
      'sequelize',
      'sequelize-typescript'
    ],
    outputFileTracingRoot: path.join(__dirname, '../../')
  },
  webpack: config => {
    config.experiments.topLevelAwait = true

    if (config.name === 'server') {
      config.optimization.concatenateModules = false
    }

    return config
  }
}

module.exports = nextConfig
