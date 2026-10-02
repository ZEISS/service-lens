import "@/config/env-config"

import { betterAuth } from "better-auth"
import { drizzleAdapter } from "better-auth/adapters/drizzle"
import { organization, twoFactor } from "better-auth/plugins"
import type { SocialProviders } from "better-auth/types"

import { db } from "@/db"
import * as schema from "@/db/schema"
import { ghec } from "@/lib/providers/ghec"

export const buildSocialProviders = () => {
  const {
    BETTER_AUTH_MICROSOFT_CLIENT_ID,
    BETTER_AUTH_MICROSOFT_CLIENT_SECRET,
    BETTER_AUTH_GOOGLE_CLIENT_ID,
    BETTER_AUTH_GOOGLE_CLIENT_SECRET,
    BETTER_AUTH_GITHUB_CLIENT_ID,
    BETTER_AUTH_GITHUB_CLIENT_SECRET,
  } = process.env

  const config: SocialProviders = {}

  if (BETTER_AUTH_GOOGLE_CLIENT_ID && BETTER_AUTH_GOOGLE_CLIENT_SECRET) {
    config.google = {
      clientId: BETTER_AUTH_GOOGLE_CLIENT_ID,
      clientSecret: BETTER_AUTH_GOOGLE_CLIENT_SECRET || "",
    }
  }

  if (BETTER_AUTH_MICROSOFT_CLIENT_ID) {
    config.microsoft = {
      clientId: BETTER_AUTH_MICROSOFT_CLIENT_ID,
      clientSecret: BETTER_AUTH_MICROSOFT_CLIENT_SECRET || "",
    }
  }

  if (BETTER_AUTH_GITHUB_CLIENT_ID && BETTER_AUTH_GITHUB_CLIENT_SECRET) {
    config.github = {
      clientId: BETTER_AUTH_GITHUB_CLIENT_ID,
      clientSecret: BETTER_AUTH_GITHUB_CLIENT_SECRET || "",
    }
  }

  return config
}

export const auth = betterAuth({
  advanced: {
    trustedProxyHeaders: true,
  },
  plugins: [
    twoFactor(),
    organization({
      teams: { enabled: true },
    }),
    ghec({
      hostName: process.env.BETTER_AUTH_GHEC_HOSTNAME,
      clientId: process.env.BETTER_AUTH_GHEC_CLIENT_ID ?? "",
      clientSecret: process.env.BETTER_AUTH_GHEC_CLIENT_SECRET ?? "",
      tokenUrl: process.env.BETTER_AUTH_GHEC_TOKEN_URL ?? "",
      authorizationUrl: process.env.BETTER_AUTH_GHEC_AUTH_URL ?? "",
    }),
  ],
  socialProviders: { ...buildSocialProviders() },
  emailAndPassword: {
    enabled: true,
    autoSignIn: true,
    disableSignUp: !(process.env.BETTER_AUTH_ENABLE_SIGNUP === "true"), // defaults to false
  },
  database: drizzleAdapter(db, {
    schema: { ...schema },
    provider: "pg",
  }),
})
