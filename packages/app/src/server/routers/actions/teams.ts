import { isAllowed, protectedProcedure } from '../../trpc'
import {
  TeamsGetSchema,
  TeamsListSchema,
  TeamsGetBySlugSchema,
  ListWorkloadByTeamSlug,
  GetTeamAndUsersByTeamSlug,
  TeamsDestroySlugSchema
} from '../schemas/teams'
import {
  findAndCountTeams,
  findOneTeam,
  findOneTeamAndMembersBySlug,
  findOneTeamBySlug,
  listWorkloadsByTeamSlug,
  destroyTeamBySlug
} from '@/db/services/teams'
import { router } from '@/server/trpc'

export const listTeams = protectedProcedure
  .input(TeamsListSchema)
  .query(async opts => await findAndCountTeams({ ...opts.input }))

export const getTeam = protectedProcedure
  .input(TeamsGetSchema)
  .query(async opts => await findOneTeam(opts.input))

export const getTeamBySlug = protectedProcedure
  .input(TeamsGetBySlugSchema)
  .query(async opts => await findOneTeamBySlug({ ...opts.input }))

export const listWorkloads = protectedProcedure
  .input(ListWorkloadByTeamSlug)
  .query(async opts => await listWorkloadsByTeamSlug({ ...opts.input }))

export const getTeamAndUsersBySlug = protectedProcedure
  .input(GetTeamAndUsersByTeamSlug)
  .query(async opts => await findOneTeamAndMembersBySlug({ ...opts.input }))

export const deleteTeamBySlug = protectedProcedure
  .use(isAllowed('write'))
  .input(TeamsDestroySlugSchema)
  .query(async opts => await destroyTeamBySlug(opts.input))

export const teamsRouter = router({
  list: listTeams,
  // add: addTeam,
  get: getTeam,
  getByName: getTeamBySlug,
  getUsersByName: getTeamAndUsersBySlug,
  listWorkloads: listWorkloads,
  deleteBySlug: deleteTeamBySlug
})
