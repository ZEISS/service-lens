import { protectedProcedure } from '../../trpc'
import {
  TeamsCreateSchema,
  TeamsGetSchema,
  TeamsListSchema
} from '../schemas/teams'
import { findAndCountTeams, findOneTeam } from '@/db/services/teams'
import { router } from '@/server/trpc'

export const listTeams = protectedProcedure
  .input(TeamsListSchema)
  .query(async opts => await findAndCountTeams({ ...opts.input }))

export const getTeam = protectedProcedure
  .input(TeamsGetSchema)
  .query(async opts => await findOneTeam(opts.input))

export const addTeam = protectedProcedure.input(TeamsCreateSchema).mutation(
  async opts => {}
  // await createTeam({ ...opts.input, userId: opts.ctx.session.user.id })
)

export const teamsRouter = router({
  list: listTeams,
  add: addTeam,
  get: getTeam
})
