'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { rhfActionDeleteWorkloadSchema } from './workload.schema'
import { destroyWorkload } from '@/db/services/workloads'
import { revalidatePath } from 'next/cache'
import { isAllowed } from '@/server/trpc'

export const rhfActionDeleteWorkload = createAction(
  protectedProcedure
    .use(isAllowed('write'))
    .input(rhfActionDeleteWorkloadSchema)
    .mutation(async opts => {
      await destroyWorkload(opts.input)
      revalidatePath('/teams/[id]/workloads', 'page')
    })
)
