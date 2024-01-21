'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createProfile } from '@/db/services/profiles'
import { isAllowed } from '@/server/trpc'

export const rhfAction = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionSchema)
    .mutation(
      async opts =>
        await createProfile({
          ...opts.input,
          scope: { resourceType: 'profile', ownerId: opts.ctx.meta.ownerId }
        })
    )
)
