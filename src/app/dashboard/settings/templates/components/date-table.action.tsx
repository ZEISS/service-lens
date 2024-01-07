'use server'

import { createAction, protectedProcedure } from '@/server/trpc'
import {
  rhfActionDeleteTemplateSchema,
  rhfActionMakeCopyTemplateSchema
} from './date-table.schema'
import {
  destroySolutionTemplate,
  makeCopySolutionTemplate
} from '@/db/services/solutions'
import { revalidatePath } from 'next/cache'

export const rhfActionDeleteTemplate = createAction(
  protectedProcedure
    .input(rhfActionDeleteTemplateSchema)
    .mutation(async opts => await destroySolutionTemplate(opts.input))
)

export const rhfActionMakeCopyTemplate = createAction(
  protectedProcedure
    .input(rhfActionMakeCopyTemplateSchema)
    .mutation(async opts => {
      await makeCopySolutionTemplate(opts.input)
      revalidatePath('/dashboard/settings/templates')
    })
)
