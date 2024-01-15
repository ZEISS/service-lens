'use server'

import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './question-form.schema'
import { addLensAnswer } from '@/db/services/workloads'

export const rhfAction = createAction(
  protectedProcedure
    .input(rhfActionSchema)
    .mutation(async opts => await addLensAnswer({ ...opts.input }))
)
