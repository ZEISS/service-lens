'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionDeleteSolutionSchema } from './solution.schema'
import { destroySolution } from '@/db/services/solutions'
import { revalidatePath } from 'next/cache'
import { isAllowed } from '@/server/trpc'

export const rhfActionDeleteSolution = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionDeleteSolutionSchema)
    .mutation(async opts => {
      await destroySolution(opts.input)
      revalidatePath('/teams/[id]/solutions', 'page')
    })
)
