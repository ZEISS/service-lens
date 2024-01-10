import { protectedProcedure } from '../../trpc'
import { ProfileListSchema, ProfileGetSchema } from '../schemas/profile'
import { router } from '@/server/trpc'

import {
  findAndCountProfiles,
  findOneProfile,
  findAllProfilesQuestions
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

export const profilesRouter = router({
  list: listProfiles,
  get: getProfile
})
