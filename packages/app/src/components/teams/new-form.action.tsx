'use server'

import 'server-only'
import { createAction, protectedProcedure } from '@/server/trpc'
import { TeamsCreateSchema } from '@/server/routers/schemas/teams'

export const rhfAction = createAction(
  protectedProcedure.input(TeamsCreateSchema).mutation(async opts => {})
)
