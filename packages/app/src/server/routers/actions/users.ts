import { protectedProcedure } from '../../trpc'
import { findOneUser, findAndCountUsers } from '@/db/services/users'
import { router } from '@/server/trpc'
import { ListUsersSchema } from '../schemas/users'

export const getUser = protectedProcedure.query(async opts =>
  findOneUser(opts.ctx.session.user.id)
)

export const listUsers = protectedProcedure
  .input(ListUsersSchema)
  .query(async opts => await findAndCountUsers({ ...opts.input }))

export const usersRouter = router({
  get: getUser,
  list: listUsers
})
