import { Team } from '@/db/models/teams'
import { User } from '@/db/models/users'
import { FindAndCountTeamsSchema, FindOneTeamSchema } from '../schemas/teams'
import type { CreateTeamSchema } from '../schemas/teams'
import { z } from 'zod'
import sequelize from '@/db/config/config'
import { UserTeam } from '../models/users-teams'

export type Pagination = {
  offset?: number
  limit?: number
}

export const createTeam = async (opts: CreateTeamSchema) =>
  sequelize.transaction(async transaction => {
    const team = await Team.create({ ...opts }, { transaction })
    await UserTeam.create(
      { userId: opts.userId, teamId: team.id },
      { transaction }
    )

    return team.dataValues
  })

export const findOneTeam = async (opts: z.infer<typeof FindOneTeamSchema>) =>
  await Team.findOne({
    where: { id: opts }
  })

export const findAndCountTeams = async (
  opts: z.infer<typeof FindAndCountTeamsSchema>
) =>
  await Team.findAndCountAll({
    order: [['name', 'DESC']],
    include: [User],
    ...opts
  })
