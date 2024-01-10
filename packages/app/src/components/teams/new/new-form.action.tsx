'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createWorkload } from '@/db/services/workloads'
import { v4 as uuidv4 } from 'uuid'

export const rhfAction = createAction(
  protectedProcedure.input(rhfActionSchema).mutation(
    async opts =>
      await createWorkload({
        id: uuidv4(),
        name: opts.input.name,
        description: opts.input.description,
        environmentsIds: opts.input.environmentsIds,
        profilesId: opts.input.profilesId,
        lensesIds: opts.input.lensesIds
      })
  )
)
