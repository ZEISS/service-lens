'use client'

import { loggerLink } from '@trpc/client'
import {
  experimental_createActionHook,
  experimental_serverActionLink
} from '@trpc/next/app-dir/client'
import { experimental_nextHttpLink } from '@trpc/next/app-dir/links/nextHttp'
import type { AppRouter } from '@/server/routers/_app'
import superjson from 'superjson'
import { getUrl } from './shared'
import { experimental_createTRPCNextAppDirClient } from './experimental'

export const api = experimental_createTRPCNextAppDirClient<AppRouter>({
  exclude() {
    return new Set(['listSolutions'])
  },
  config() {
    return {
      transformer: superjson,
      links: [
        loggerLink({
          enabled: op => true
        }),
        experimental_nextHttpLink({
          batch: true,
          url: getUrl(),
          headers() {
            return {
              'x-trpc-source': 'client'
            }
          }
        })
      ]
    }
  }
})

export const useAction = experimental_createActionHook({
  links: [loggerLink({ enabled: op => true }), experimental_serverActionLink()],
  transformer: superjson
})
