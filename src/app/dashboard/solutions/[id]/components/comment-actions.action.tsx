'use server'

import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfDeleteCommentActionSchema } from './comment-actions.schema'
import { deleteSolutionComment } from '@/db/services/solutions'
import { revalidatePath } from 'next/cache'

export const rhfDeleteCommentAction = createAction(
  protectedProcedure
    .input(rhfDeleteCommentActionSchema)
    .mutation(async opts => {
      await deleteSolutionComment(opts.input)
      revalidatePath('/dashboard/solutions/[id]', 'page')
    })
)
