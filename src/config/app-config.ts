import packageJson from "../../package.json"

const currentYear = new Date().getFullYear()

export const APP_CONFIG = {
  name: "Service Lens",
  version: packageJson.version,
  copyright: `© ${currentYear}, Service Lens.`,
  meta: {
    title: "Service Lens - A modern-day service governance platform",
    description:
      "Service Lens is a modern-day service governance platform that helps you manage, monitor, and optimize your services with ease.",
  },
  enabledSocialProviders: {
    microsoft: process.env.BETTER_AUTH_MICROSOFT_ENABLED === "true",
    github: process.env.BETTER_AUTH_GITHUB_ENABLED === "true",
    google: process.env.BETTER_AUTH_GOOGLE_ENABLED === "true",
  },
}
