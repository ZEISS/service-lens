import { protectedProcedure } from '../../trpc'
import { EnvironmentListSchema } from '../schemas/environments'
import { findAndCountEnvironments } from '@/db/services/environments'

export const listEnvironments = protectedProcedure
  .input(EnvironmentListSchema)
  .query(async opts => await findAndCountEnvironments({ ...opts.input }))
