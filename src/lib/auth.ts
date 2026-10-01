import "@/config/env-config"

import { betterAuth } from "better-auth"
import { drizzleAdapter } from "better-auth/adapters/drizzle"
import { organization, twoFactor } from "better-auth/plugins"

import { db } from "@/db"
import * as schema from "@/db/schema"
import { ghec } from "@/lib/providers/ghec"

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
    autoSignIn: true,
    disableSignUp: !(process.env.BETTER_AUTH_ENABLE_SIGNUP === "true"), // defaults to false
  },
  database: drizzleAdapter(db, {
    schema: { ...schema },
    provider: "pg",
  }),
})
