'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { createTeam } from '@/db/services/teams'
import { TeamsCreateSchema } from '@/server/routers/schemas/teams'

export const rhfAction = createAction(
  protectedProcedure
    .input(TeamsCreateSchema)
    .mutation(
      async opts =>
        await createTeam({ ...opts.input, userId: opts.ctx.session.user.id })
    )
)
