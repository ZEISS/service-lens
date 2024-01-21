import { protectedProcedure } from '../../trpc'
import {
  TeamsGetSchema,
  TeamsListSchema,
  TeamsGetBySlugSchema,
  ListWorkloadByTeamSlug
} from '../schemas/teams'
import {
  findAndCountTeams,
  findOneTeam,
  findOneTeamBySlug,
  listWorkloadsByTeamSlug
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

export const teamsRouter = router({
  list: listTeams,
  // add: addTeam,
  get: getTeam,
  getByName: getTeamBySlug,
  listWorkloads: listWorkloads
})
