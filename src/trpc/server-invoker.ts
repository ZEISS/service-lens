import { appRouter } from '@/server/routers/_app'
import { auth } from '@/auth'
import { cookies } from 'next/headers'
import { experimental_createTRPCNextAppDirServer } from '@trpc/next/app-dir/server'
import { experimental_nextCacheLink } from '@trpc/next/app-dir/links/nextCache'
import { loggerLink } from '@trpc/client'
import { type inferRouterInputs, type inferRouterOutputs } from '@trpc/server'
import SuperJSON from 'superjson'
import type { AppRouter } from '@/server/routers/_app'

export const api = experimental_createTRPCNextAppDirServer<typeof appRouter>({
  config() {
    return {
      transformer: SuperJSON,
      links: [
        loggerLink({
          enabled: opts =>
            process.env.NODE_ENV === 'development' ||
            (opts.direction === 'down' && opts.result instanceof Error)
        }),
        experimental_nextCacheLink({
          // requests are cached for 5 seconds
          revalidate: 5,
          router: appRouter,
          createContext: async () => {
            return {
              session: await auth(),
              headers: {
                cookie: cookies().toString(),
                'x-trpc-source': 'rsc-invoke'
              }
            }
          }
        })
      ]
    }
  }
})

export type RouterInputs = inferRouterInputs<AppRouter>
export type RouterOutputs = inferRouterOutputs<AppRouter>
