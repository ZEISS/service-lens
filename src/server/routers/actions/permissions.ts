import { PermissionGetSchema } from '../schemas/permissions'
import { protectedProcedure } from '../../trpc'
import { findOnePermission } from '@/db/services/permissions'

export const checkPermission = protectedProcedure
  .input(PermissionGetSchema)
  .query(async opts => findOnePermission(opts.input))
