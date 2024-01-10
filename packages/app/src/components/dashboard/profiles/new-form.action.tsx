'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { createProfile } from '@/db/services/profiles'

export const rhfAction = createAction(
  protectedProcedure
    .input(rhfActionSchema)
    .mutation(async opts => await createProfile({ ...opts.input }))
)
