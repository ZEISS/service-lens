import {
  clientCallTypeToProcedureType,
  CreateTRPCProxyClient,
  createTRPCUntypedClient
} from '@trpc/client'
import { LRUCache } from 'lru-cache'
import { AnyRouter } from '@trpc/server'
import { createRecursiveProxy } from '@trpc/server/shared'
import { CreateTRPCClientOptions } from '@trpc/client'

export type Op = String
export type Ops = Set<Op>

export interface CreateTRPCNextAppRouterOptions<TRouter extends AnyRouter> {
  config: () => CreateTRPCClientOptions<TRouter>
  exclude: () => Ops
}

const options = {
  max: 100,
  ttl: 1000 * 60 * 5
}

export function experimental_createTRPCNextAppDirClient<
  TRouter extends AnyRouter
>(opts: CreateTRPCNextAppRouterOptions<TRouter>) {
  const client = createTRPCUntypedClient<TRouter>(opts.config())
  const cache = new LRUCache<String, Promise<unknown>>(options)
  const exclude = opts?.exclude()

  return createRecursiveProxy(({ path, args }) => {
    const pathCopy = [...path]
    const procedureType = clientCallTypeToProcedureType(pathCopy.pop()!)
    const fullPath = pathCopy.join('.')

    if (procedureType === 'query' && !exclude?.has(fullPath)) {
      const queryCacheKey = JSON.stringify([path, args[0]])
      const cached = cache.get(queryCacheKey)

      if (cached) {
        return cached
      }
    }

    const promise: Promise<unknown> = (client as any)[procedureType](
      fullPath,
      ...args
    )
    if (procedureType !== 'query') {
      return promise
    }

    const queryCacheKey = JSON.stringify([path, args[0]])

    cache.set(queryCacheKey, promise)

    return promise
  }) as CreateTRPCProxyClient<TRouter>
}
