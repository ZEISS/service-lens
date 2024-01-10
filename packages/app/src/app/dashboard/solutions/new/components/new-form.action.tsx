'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionSchema } from './new-form.schema'
import { addSolution } from '@/db/services/solutions'
import { v4 as uuidv4 } from 'uuid'
import { User } from '@/db/models/users'

export const rhfAction = createAction(
  protectedProcedure.input(rhfActionSchema).mutation(async opts => {
    const user = await User.findOne({
      where: { email: opts.ctx.session.user.email ?? '' }
    })

    return await addSolution({
      id: uuidv4(),
      userId: user?.id,
      ...opts.input
    })
  })
)
