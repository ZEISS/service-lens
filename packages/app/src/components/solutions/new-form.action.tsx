'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionNewSolutionSchema } from './new-form.schema'
import { createSolution } from '@/db/services/solutions'
import { isAllowed } from '@/server/trpc'

export const rhfActionCreateNewSolution = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionNewSolutionSchema)
    .mutation(
      async opts =>
        await createSolution({
          ...opts.input,
          userId: opts.ctx.session?.user.id ?? opts.ctx.meta.ownerId,
          ownerId: opts.ctx.meta.ownerId,
          resourceType: 'solution'
        })
    )
)
