import { experimental_createServerActionHandler } from '@trpc/next/app-dir/server'
import { initTRPC, TRPCError } from '@trpc/server'
import { auth } from '@/auth'
import { headers } from 'next/headers'
import superjson from 'superjson'
import { ZodError } from 'zod'
import type { Context } from './context'
import { type Permissions } from '@/db/models/permissions'
import cookie from 'cookie'
import { findOnePermission } from '@/db/services/permissions'
import { findOneTeamBySlug } from '@/db/services/teams'

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
    const { headers } = ctx
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
      userId: ctx.session?.user.id ?? '',
      permission: 'read'
    })

    if (!allowed) {
      throw new TRPCError({
        code: 'FORBIDDEN'
      })
    }

    return next({ ctx })
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
