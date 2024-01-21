import { auth } from '@/auth'
import { experimental_createServerActionHandler } from '@trpc/next/app-dir/server'
import { findOnePermission } from '@/db/services/permissions'
import { findOneTeamBySlug } from '@/db/services/teams'
import { headers } from 'next/headers'
import { initTRPC, TRPCError } from '@trpc/server'
import { type Permissions } from '@/db/models/permissions'
import { ZodError } from 'zod'
import cookie from 'cookie'
import superjson from 'superjson'
import type { Context } from './context'

const t = initTRPC.context<Context>().create({
  transformer: superjson,
  errorFormatter({ shape, error }) {
    return {
      ...shape,
      data: {
        ...shape.data,
        zodError:
          error.code === 'BAD_REQUEST' && error.cause instanceof ZodError
            ? error.cause.flatten()
            : null
      }
    }
  }
})

export const router = t.router
export const publicProcedure = t.procedure

const isAuthenticated = t.middleware(async ({ ctx, next }) => {
  const { session } = ctx

  if (!session?.user) {
    throw new TRPCError({
      code: 'UNAUTHORIZED'
    })
  }

  return next({ ctx: { session } })
})
export const protectedProcedure = publicProcedure.use(isAuthenticated)

export const isAllowed = (permission: Permissions) =>
  t.middleware(async ({ ctx, next }) => {
    const { headers, session } = ctx
    const hasCookie = headers && 'cookie' in headers

    if (!hasCookie) {
      throw new TRPCError({
        code: 'FORBIDDEN'
      })
    }

    const cookies = new Map(Object.entries(cookie.parse(headers['cookie'])))
    const hasScope = cookies.has('scope')
    const scope = cookies.get('scope')

    if (!hasScope) {
      throw new TRPCError({
        code: 'FORBIDDEN'
      })
    }

    const team = await findOneTeamBySlug({ slug: scope ?? '' })

    if (!team) {
      throw new TRPCError({
        code: 'FORBIDDEN'
      })
    }

    const allowed = await findOnePermission({
      teamId: team.id,
      userId: session?.user.id ?? '',
      permission
    })

    if (!allowed) {
      throw new TRPCError({
        code: 'FORBIDDEN'
      })
    }

    return next({ ctx: { ...ctx, meta: { ownerId: team.id } } })
  })

export const createAction = experimental_createServerActionHandler(t, {
  async createContext() {
    const session = await auth()

    return {
      session,
      headers: {
        cookie: headers().get('cookie') ?? ''
      }
    }
  }
})
