'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { destroyTeam } from '@/db/services/teams'
import { isAllowed } from '@/server/trpc'
import { cookies } from 'next/headers'

export const rhfActionDeleteTeam = createAction(
  protectedProcedure.use(isAllowed('write')).mutation(async opts => {
    await destroyTeam(opts.ctx.meta.ownerId)
    cookies().set('scope', 'personal')
  })cd
)
