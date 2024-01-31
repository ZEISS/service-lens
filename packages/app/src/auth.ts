import NextAuth from 'next-auth'
import GithubProvider from 'next-auth/providers/github'
import SequelizeAdapter from '@/lib/adapter-sequelize'
import { type DefaultSession, NextAuthConfig } from 'next-auth'
import sequelize from '@/db/config/config'

const env = process.env.NODE_ENV || 'development'
const isProduction = env === 'production'

const adapter = SequelizeAdapter(sequelize, {
  synchronize: false
})

declare module 'next-auth' {
  interface Session {
    user: DefaultSession['user'] & {
      id: string
    }
  }
}

export const options = {
  providers: [
    GithubProvider({
      clientId: process.env.GITHUB_ID!,
      clientSecret: process.env.GITHUB_SECRET!
    })
  ],
  session: {
    generateSessionToken: () => crypto.randomUUID()
  },
  adapter,
  debug: !isProduction,
  pages: {
    signIn: '/login'
  },
  callbacks: {
    session: async ({ session, user }: any) => {
      if (user?.id) {
        session.user.id = user.id
      }

      return session
    }
  }
} satisfies NextAuthConfig

export const {
  handlers: { GET, POST },
  auth
} = NextAuth(options)
