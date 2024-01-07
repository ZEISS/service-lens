'use server'

import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './comment-form.schema'
import { User } from '@/db/models/users'
import { addSolutionComment } from '@/db/services/solutions'
import { revalidatePath } from 'next/cache'

export const rhfAction = createAction(
  protectedProcedure.input(rhfActionSchema).mutation(async opts => {
    const user = await User.findOne({
      where: { email: opts.ctx.session.user?.email ?? '' }
    })

    const comment = await addSolutionComment({
      ...opts.input,
      userId: user!.id
    })

    revalidatePath('/dashboard/solutions/[id]', 'page')

    return comment
  })
)
