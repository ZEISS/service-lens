# Service Lens :eyeglasses:

[![Test & Build](https://github.com/zeiss/service-lens/actions/workflows/main.yml/badge.svg)](https://github.com/zeiss/service-lens/actions/workflows/main.yml)
[![Taylor Swift](https://img.shields.io/badge/secured%20by-taylor%20swift-brightgreen.svg)](https://twitter.com/SwiftOnSecurity)
[![Volkswagen](https://auchenberg.github.io/volkswagen/volkswargen_ci.svg?v=1)](https://github.com/auchenberg/volkswagen)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

> :warning: This project is in early development. It is not ready for production use.

## About

Service Lens is an enterprise service management tool. It allows you to manage your services, identify risks, review the solutions that created them and the business context of every service. It follows the Well-Architected methodology established by AWS and Microsoft.

![preview](assets/screenshot_1.png)

It is build on [fiber-htmx](https://github.com/ZEISS/fiber-htmx) and uses a 3-tier architecture.

## Development

Please, set all environment variables in `.env`. `docker compose up db` will launch a local development database.

```
air
```

This launches a development instance of the application.

# License

[LICENSE](./LICENSE)
