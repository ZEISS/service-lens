import "@/config/env-config"

import { betterAuth } from "better-auth"
import { drizzleAdapter } from "better-auth/adapters/drizzle"
import { organization, twoFactor } from "better-auth/plugins"
import { genericOAuth } from "better-auth/plugins"

import { db } from "@/db"
import * as schema from "@/db/schema"

export const auth = betterAuth({
  advanced: {
    trustedProxyHeaders: true,
  },
  plugins: [
    twoFactor(),
    organization({
      teams: { enabled: true },
    }),
    genericOAuth({
      config: [
        {
          providerId: "ghec",
          name: "GitHub Enterprise Cloud",
          clientId: process.env.BETTER_AUTH_GENERIC_OAUTH_CLIENT_ID ?? "",
          clientSecret: process.env.BETTER_AUTH_GENERIC_OAUTH_CLIENT_SECRET ?? "",
          tokenUrl: process.env.BETTER_AUTH_GENERIC_OAUTH_TOKEN_URL ?? "",
          authorizationUrl: process.env.BETTER_AUTH_GENERIC_OAUTH_AUTH_URL ?? "",
          tokenEndpointAuth: {
            method: "client_secret_post",
          },
          scopes: ["email", "read:user"]
        },
      ],
    }),
  ],
  socialProviders: {
    microsoft: {
      clientId: process.env.BETTER_AUTH_MICROSOFT_CLIENT_ID || "",
      clientSecret: process.env.BETTER_AUTH_MICROSOFT_CLIENT_SECRET || "",
    },
    google: {
      clientId: process.env.BETTER_AUTH_GOOGLE_CLIENT_ID || "",
      clientSecret: process.env.BETTER_AUTH_GOOGLE_CLIENT_SECRET || "",
    },
    github: {
      clientId: process.env.BETTER_AUTH_GITHUB_CLIENT_ID || "",
      clientSecret: process.env.BETTER_AUTH_GITHUB_CLIENT_SECRET || "",
    },
  },
  emailAndPassword: {
    enabled: true,
    autoSignIn: true, //defaults to true
  },
  database: drizzleAdapter(db, {
    schema: { ...schema },
    provider: "pg",
  }),
})
