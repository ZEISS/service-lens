# Service Lens :eyeglasses:

[![Test & Build](https://github.com/katallaxie/service-lens/actions/workflows/main.yml/badge.svg)](https://github.com/katallaxie/service-lens/actions/workflows/main.yml)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> :warning: This project is in early development. It is not ready for production use.

## About

Service Lens is an enterprise service management tool. It allows you to manage your services, identify risks, review the solutions that created them and the business context of every service. It follows the Well-Architected methodology established by AWS and Microsoft.

## Development

Please create a `.env.local` file from the `.env.example` file. You need to [create a OAuth app](https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/creating-an-oauth-app) on GitHub to get the `GITHUB_ID` and `GITHUB_SECRET`.

```bash
# Install dependencies
npm i

# Run Postgres
docker compose up

# Run migrations
npm run migrate:up

# Run seeds
npm run db:seed

# Run the server
npm run dev
```

# License

[LICENSE](./LICENSE)