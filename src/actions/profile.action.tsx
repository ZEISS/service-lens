'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionDeleteProfileSchema } from './profile.schema'
import { destroyProfile } from '@/db/services/profiles'
import { revalidatePath } from 'next/cache'

export const rhfActionDeleteProfile = createAction(
  protectedProcedure
    .input(rhfActionDeleteProfileSchema)
    .mutation(async opts => {
      await destroyProfile(opts.input)
      revalidatePath('/dashboard/profiles')
    })
)
