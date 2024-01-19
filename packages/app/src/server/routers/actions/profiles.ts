import { protectedProcedure } from '../../trpc'
import {
  ProfileListSchema,
  ProfileGetSchema,
  ListProfileByTeamSlug
} from '../schemas/profile'
import { router } from '@/server/trpc'

import {
  findAndCountProfiles,
  findOneProfile,
  findAllProfilesQuestions,
  listProfileByTeamSlug
} from '@/db/services/profiles'

export const listProfiles = protectedProcedure
  .input(ProfileListSchema)
  .query(async opts => await findAndCountProfiles({ ...opts.input }))
export const getProfile = protectedProcedure
  .input(ProfileGetSchema)
  .query(async opts => await findOneProfile(opts.input))
export const listProfilesQuestions = protectedProcedure.query(
  async opts => await findAllProfilesQuestions(opts)
)

export const listByTeam = protectedProcedure
  .input(ListProfileByTeamSlug)
  .query(async opts => await listProfileByTeamSlug({ ...opts.input }))

export const profilesRouter = router({
  list: listProfiles,
  listByTeam: listByTeam,
  get: getProfile
})
