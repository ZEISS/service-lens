'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createLens } from '@/db/services/lenses'
import { isAllowed } from '@/server/trpc'

export const rhfAction = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionSchema)
    .mutation(
      async opts =>
        await createLens({
          name: opts.input.name,
          description: opts.input.description,
          spec: opts.input.spec,
          ownerId: opts.ctx.meta.ownerId,
          resourceType: 'lens'
        })
    )
)
