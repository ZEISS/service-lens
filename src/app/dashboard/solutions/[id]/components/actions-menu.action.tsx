'use server'

import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfDeleteSolutionActionSchema } from './actions-menu.schema'
import { deleteSolution } from '@/db/services/solutions'
import { revalidatePath } from 'next/cache'

export const rhfDeleteSolutionAction = createAction(
  protectedProcedure
    .input(rhfDeleteSolutionActionSchema)
    .mutation(async opts => {
      await deleteSolution(opts.input)
      revalidatePath('/dashboard/solutions')
    })
)
