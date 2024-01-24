'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionDeleteTeamSchema } from './settings-teams.schema'
import { destroyTeam } from '@/db/services/teams'
import { revalidatePath } from 'next/cache'

export const rhfActionDeleteTeam = createAction(
  protectedProcedure.input(rhfActionDeleteTeamSchema).mutation(async opts => {
    await destroyTeam(opts.input)
    revalidatePath('/settings/teams')
  })
)
