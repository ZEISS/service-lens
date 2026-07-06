import { organizationClient } from "better-auth/client/plugins"
import { createAuthClient } from "better-auth/react"

export const client = createAuthClient({
  /** The base URL of the server (optional if you're using the same domain) */
  baseURL: process.env.BETTER_AUTH_URL ?? "http://localhost:3000",
  plugins: [
    organizationClient({
      teams: { enabled: true },
    }),
  ],
})

export const { signIn, signUp, signOut, useSession, useListOrganizations, useActiveOrganization } = client
