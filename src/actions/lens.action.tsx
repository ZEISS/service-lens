'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import {
  rhfActionDeleteLensSchema,
  rhfActionPublishLensSchema
} from './lens.schema'
import { deleteLens, publishLens } from '@/db/services/lenses'
import { revalidatePath } from 'next/cache'
import { isAllowed } from '@/server/trpc'

export const rhfActionDeleteLens = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionDeleteLensSchema)
    .mutation(async opts => {
      await deleteLens({
        lensId: opts.input,
        resourceType: 'lens',
        ownerId: opts.ctx.meta.ownerId
      })
      revalidatePath('/teams/[id]/lenses', 'page')
    })
)

export const rhfActionPushlishLens = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionPublishLensSchema)
    .mutation(async opts => {
      await publishLens({
        lensId: opts.input,
        resourceType: 'lens',
        ownerId: opts.ctx.meta.ownerId
      })
      revalidatePath('/teams/[id]/lenses', 'page')
    })
)
