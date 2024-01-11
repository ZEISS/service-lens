'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createWorkload } from '@/db/services/workloads'

export const rhfAction = createAction(
  protectedProcedure.input(rhfActionSchema).mutation(
    async opts =>
      await createWorkload({
        name: opts.input.name,
        description: opts.input.description,
        environmentsIds: opts.input.environmentsIds,
        profilesId: opts.input.profilesId,
        lensesIds: opts.input.lensesIds
      })
  )
)
