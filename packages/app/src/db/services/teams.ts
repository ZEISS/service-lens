import { Team } from '@/db/models/team'
import { User } from '@/db/models/users'
import {
  FindAndCountTeamsSchema,
  FindOneTeamSchema,
  CreateTeamSchema
} from '../schemas/teams'
import { z } from 'zod'
import sequelize from '@/db/config/config'

export type Pagination = {
  offset?: number
  limit?: number
}

export const createTeam = async (opts: z.infer<typeof CreateTeamSchema>) =>
  sequelize.transaction(
    async transaction => await Team.create({ ...opts }, { transaction })
  )

export const findOneTeam = async (opts: z.infer<typeof FindOneTeamSchema>) =>
  await Team.findOne({
    where: { id: opts },
    include: [User]
  })

export const findAndCountTeams = async (
  opts: z.infer<typeof FindAndCountTeamsSchema>
) =>
  await Team.findAndCountAll({
    order: [['name', 'DESC']],
    include: [User],
    ...opts
  })
