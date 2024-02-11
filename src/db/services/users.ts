import { Team } from '@/db/models/teams'
import { User } from '@/db/models/users'
import {
  type FindOneUserSchema,
  type FindAndCountUsers,
  type DestroyUser
} from '../schemas/users'
import sequelize from '@/db/config/config'
import { Workload } from '@/db/models/workload'

export const findOneUser = async (opts: FindOneUserSchema) =>
  await User.findOne({
    where: { id: opts },
    include: [Team, Workload]
  })

export const destroyUser = async (opts: DestroyUser) =>
  sequelize.transaction(
    async transaction =>
      await User.destroy({ where: { id: opts }, transaction })
  )

export const findAndCountUsers = async (opts: FindAndCountUsers) =>
  await User.findAndCountAll({
    order: [['name', 'DESC']],
    ...opts
  })
