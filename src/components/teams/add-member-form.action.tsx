'use server'

import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionAddMemberSchema } from './add-member-form.schema'
import { User } from '@/db/models/users'
import { addSolutionComment } from '@/db/services/solutions'
import { revalidatePath } from 'next/cache'

export const rhfActionAddMember = createAction(
  protectedProcedure.input(rhfActionAddMemberSchema).mutation(async opts => {
    //  revalidatePath('/dashboard/solutions/[id]', 'page')
  })
)
