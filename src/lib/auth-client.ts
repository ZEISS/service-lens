import { organizationClient } from "better-auth/client/plugins"
import { createAuthClient } from "better-auth/react"

export const client = createAuthClient({
  plugins: [
    organizationClient({
      teams: { enabled: true },
    }),
  ],
})

export const { signIn, signUp, signOut, useSession, useListOrganizations, useActiveOrganization } = client
