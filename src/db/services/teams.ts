import { Team } from '@/db/models/teams'
import { User } from '@/db/models/users'
import {
  FindAndCountTeamsSchema,
  FindOneTeamSchema,
  type FineOneTeamAndUsersByNameSlug,
  type CreateTeamSchema
} from '../schemas/teams'
import type {
  FindOneTeamByNameSlug,
  ListWorkloadsByTeamSlug,
  DestroyTeam
} from '../schemas/teams'
import { z } from 'zod'
import sequelize from '@/db/config/config'
import { UserTeam } from '../models/users-teams'
import { Workload } from '../models/workload'
import { UserRole } from '../models/users-roles'

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
    await UserRole.create(
      {
        userId: opts.userId,
        teamId: team.id,
        roleId: BigInt(1)
      },
      { transaction }
    )

    return team.dataValues
  })

export const destroyTeam = async (opts: DestroyTeam) =>
  sequelize.transaction(
    async transaction =>
      await Team.destroy({ where: { id: opts }, transaction })
  )

export const findOneTeam = async (opts: z.infer<typeof FindOneTeamSchema>) =>
  await Team.findOne({
    where: { id: opts }
  })

export const findOneTeamBySlug = async (opts: FindOneTeamByNameSlug) =>
  await Team.findOne({
    where: { slug: opts.slug }
  })

export const findOneTeamAndMembersBySlug = async (
  opts: FineOneTeamAndUsersByNameSlug
) =>
  await Team.findOne({
    where: { slug: opts.slug },
    include: [{ model: User }]
  })

export const listWorkloadsByTeamSlug = async (opts: ListWorkloadsByTeamSlug) =>
  await Team.findOne({
    where: { slug: opts.slug },
    include: [{ model: Workload }]
  })

export const findAndCountTeams = async (
  opts: z.infer<typeof FindAndCountTeamsSchema>
) =>
  await Team.findAndCountAll({
    order: [['name', 'DESC']],
    include: [User],
    ...opts
  })
