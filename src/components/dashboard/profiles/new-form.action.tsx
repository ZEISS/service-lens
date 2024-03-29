'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createProfile } from '@/db/services/profiles'
import { isAllowed } from '@/server/trpc'
import { revalidatePath } from 'next/cache'

export const rhfAction = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionSchema)
    .mutation(async opts => {
      revalidatePath('/teams/[id]/lenses', 'page')
      return await createProfile({
        ...opts.input,
        scope: { resourceType: 'profile', ownerId: opts.ctx.meta.ownerId }
      })
    })
)
