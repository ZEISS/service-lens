'use server'

import 'server-only'
import { createAction, isAllowed, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createWorkload } from '@/db/services/workloads'

export const rhfAction = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionSchema)
    .mutation(
      async opts =>
        await createWorkload({
          name: opts.input.name,
          description: opts.input.description,
          ownerId: opts.ctx.meta.ownerId,
          resourceType: 'workload'
        })
    )
)
