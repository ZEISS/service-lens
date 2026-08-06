import { db } from "@/db"
import * as schema from "@/db/schema"
import { betterAuth } from "better-auth"
import { drizzleAdapter } from "better-auth/adapters/drizzle"
import { organization } from "better-auth/plugins"
import { twoFactor } from "better-auth/plugins"

export const auth = betterAuth({
  plugins: [
    twoFactor(),
    organization({
      teams: { enabled: true },
    }),
  ],
  socialProviders: {
    microsoft: {
      clientId: process.env.BETTER_AUTH_MICROSOFT_CLIENT_ID || "",
      clientSecret: process.env.BETTER_AUTH_MICROSOFT_CLIENT_SECRET || "",
    },
    github: {
      clientId: process.env.BETTER_AUTH_GITHUB_CLIENT_ID || "",
      clientSecret: process.env.BETTER_AUTH_GITHUB_CLIENT_SECRET || "",
    },
  },
  baseURL: process.env.NEXT_PUBLIC_BASE_URL ?? "http://localhost:3000",
  emailAndPassword: {
    enabled: true,
  },
  database: drizzleAdapter(db, {
    schema: { ...schema },
    provider: "pg",
  }),
})
