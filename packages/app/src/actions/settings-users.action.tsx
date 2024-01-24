'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionDeleteUserSchema } from './settings-users.schema'
import { destroyUser } from '@/db/services/users'
import { revalidatePath } from 'next/cache'

export const rhfActionDeleteTeam = createAction(
  protectedProcedure.input(rhfActionDeleteUserSchema).mutation(async opts => {
    await destroyUser(opts.input)
    revalidatePath('/settings/users')
  })
)
