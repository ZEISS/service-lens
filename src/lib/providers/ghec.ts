import { genericOAuth } from "better-auth/plugins"
import { betterFetch } from "@better-fetch/fetch"
import type { GenericOAuthUserInfo } from "better-auth/plugins/generic-oauth"
import type { GithubProfile } from "better-auth"

export type GHECOptions = {
  hostName?: string
  clientId: string
  clientSecret: string
  tokenUrl: string
  authorizationUrl: string
}

export const ghec = ({ hostName, clientId, clientSecret, tokenUrl, authorizationUrl }: GHECOptions) => genericOAuth({
  config: [
    {
      providerId: "ghec",
      name: "GitHub Enterprise Cloud",
      clientId,
      clientSecret,
      tokenUrl,
      authorizationUrl,
      getUserInfo: async (token) => {
        const { data: profile, error } = await betterFetch<GithubProfile>(
              `https://api.${hostName}/user`,
              {
                headers: {
                  "User-Agent": "better-auth",
                  Authorization: `token ${token?.raw?.["access_token"]}`,
                },
              },
            );

            if (error || !profile) {
              return null;
            }

            return {
              id: String(profile.id),
              name: profile.name || profile.login,
              email: profile.email,
              image: profile.avatar_url,
            } as GenericOAuthUserInfo;
      },
      scopes: ["read:user", "user:email"]
    },
  ],
})
