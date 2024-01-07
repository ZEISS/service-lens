import { v4 as uuidv4 } from 'uuid'
import { Profile } from '@/db/models/profile'
import { ProfileQuestion } from '@/db/models/profile-question'
import { ProfileQuestionAnswer } from '@/db/models/profile-question-answers'
import { ProfileQuestionChoice } from '@/db/models/profile-question-choice'
import {
  FindAndCountProfilesSchema,
  FindOneProfileSchema,
  CreateProfileSchema,
  FindAllProfilesQuestionsSchema,
  DestroyProfileSchema
} from '../schemas/profiles'
import { z } from 'zod'
import sequelize from '@/db/config/config'

export type Pagination = {
  offset?: number
  limit?: number
}

export const createProfile = async (
  opts: z.infer<typeof CreateProfileSchema>
) =>
  sequelize.transaction(async transaction => {
    const profile = await Profile.create(
      { id: uuidv4(), ...opts },
      { transaction }
    )

    await ProfileQuestionAnswer.bulkCreate(
      Object.values(opts.selectedChoices)
        .flatMap(choices => choices)
        .map(choice => ({
          choiceId: BigInt(choice),
          profileId: profile.id
        })),
      { transaction }
    )

    return profile
  })

export const destroyProfile = async (
  opts: z.infer<typeof DestroyProfileSchema>
) =>
  sequelize.transaction(
    async transaction =>
      await Profile.destroy({ where: { id: opts }, transaction })
  )

// export async function deleteProfile(id: string) {
//   return await Profile.destroy({ where: { id } })
// }

export const findOneProfile = async (
  opts: z.infer<typeof FindOneProfileSchema>
) =>
  await Profile.findOne({
    where: { id: opts },
    include: [{ model: ProfileQuestionChoice, include: [ProfileQuestion] }]
  })

export const findAndCountProfiles = async (
  opts: z.infer<typeof FindAndCountProfilesSchema>
) =>
  await Profile.findAndCountAll({
    order: [['name', 'DESC']],
    include: [ProfileQuestionChoice],
    ...opts
  })

export const findAllProfilesQuestions = async (
  opts: z.infer<typeof FindAllProfilesQuestionsSchema>
) =>
  await ProfileQuestion.findAll({
    order: [['name', 'DESC']],
    include: [ProfileQuestionChoice],
    ...opts
  })
