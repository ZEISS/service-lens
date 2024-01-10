'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import {
  rhfActionDeleteLensSchema,
  rhfActionPublishLensSchema
} from './lens.schema'
import { deleteLens, publishLens } from '@/db/services/lenses'
import { revalidatePath } from 'next/cache'

export const rhfActionDeleteLens = createAction(
  protectedProcedure.input(rhfActionDeleteLensSchema).mutation(async opts => {
    await deleteLens(opts.input)
    revalidatePath('/dashboard/lenses')
  })
)

export const rhfActionPushlishLens = createAction(
  protectedProcedure.input(rhfActionPublishLensSchema).mutation(async opts => {
    await publishLens(opts.input)
    revalidatePath(`/dashboard/lenses/${opts.input}`)
  })
)
