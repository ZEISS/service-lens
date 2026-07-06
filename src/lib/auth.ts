import { db } from "@/db"
import * as schema from "@/db/schema"
import { betterAuth } from "better-auth"
import { drizzleAdapter } from "better-auth/adapters/drizzle"
import { organization } from "better-auth/plugins"

export const auth = betterAuth({
  plugins: [
    organization({
      teams: { enabled: true },
    }),
  ],
  baseURL: process.env.NEXT_PUBLIC_BASE_URL ?? "http://localhost:3000",
  emailAndPassword: {
    enabled: true,
  },
  database: drizzleAdapter(db, {
    schema: { ...schema },
    provider: "pg",
  }),
})
