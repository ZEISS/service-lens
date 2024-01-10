import { fetchRequestHandler } from '@trpc/server/adapters/fetch'
import { createContext } from '@/server/context'
import { appRouter } from '@/server/routers/_app'

// Add back once NextAuth v5 is released
// export const runtime = 'edge';

const handler = (req: Request) =>
  fetchRequestHandler({
    endpoint: '/api/trpc',
    req,
    router: appRouter,
    createContext,
    onError:
      process.env.NODE_ENV === 'development'
        ? ({ path, error }) => {
            console.error(
              `❌ tRPC failed on ${path ?? '<no-path>'}: ${error.message}`
            )
          }
        : undefined
  })

export { handler as GET, handler as POST }
