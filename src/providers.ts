import GithubProvider from 'next-auth/providers/github'

export const providers = [
  GithubProvider({
    clientId: process.env.GITHUB_ID!,
    clientSecret: process.env.GITHUB_SECRET!
  })
]

export default providers
