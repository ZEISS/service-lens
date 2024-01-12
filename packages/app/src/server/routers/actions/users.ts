import { protectedProcedure } from '../../trpc'
import { findOneUser } from '@/db/services/users'
import { router } from '@/server/trpc'

export const getUser = protectedProcedure.query(async opts =>
  findOneUser(opts.ctx.session.user.id)
)

export const usersRouter = router({
  get: getUser
})
