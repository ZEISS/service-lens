import { v4 as uuidv4 } from 'uuid'
import { Lens } from '@/db/models/lens'
import { LensPillar } from '@/db/models/lens-pillars'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { LensPillarQuestionResource } from '@/db/models/lens-pillar-questions-resources'
import { Spec } from '@/db/schemas/spec'
import {
  LensesGetSchema,
  LensesDeleteSchema,
  LensesPublishSchema,
  LensesGetQuestionSchema
} from '../schemas/lenses'
import sequelize from '@/db/config/config'
import { z } from 'zod'
import { LensPillarResource } from '../models/lens-pillar-resources'

export type Pagination = {
  offset?: number
  limit?: number
}

export const getLens = async (opts: z.infer<typeof LensesGetSchema>) =>
  await Lens.findOne({
    where: { id: opts },
    include: [{ model: LensPillar, include: [{ model: LensPillarQuestion }] }]
  })

export const deleteLens = async (opts: z.infer<typeof LensesDeleteSchema>) =>
  await Lens.destroy({
    where: { id: opts }
  })

export const publishLens = async (opts: z.infer<typeof LensesPublishSchema>) =>
  await Lens.update({ isDraft: false }, { where: { id: opts } })

export async function createLens({
  id = uuidv4(),
  name,
  description,
  spec
}: {
  id: string
  name: string
  description: string
  spec: string
}) {
  return await sequelize.transaction(async transaction => {
    const s = await Spec.parseAsync(JSON.parse(spec))

    const lens = await Lens.create(
      {
        id,
        name,
        version: s.version,
        description,
        spec: s,
        isDraft: true
      },
      {
        transaction
      }
    )

    const pillars = await LensPillar.bulkCreate(
      [
        ...s.pillars.map(pillar => ({
          lensId: lens.id,
          name: pillar.name,
          ref: pillar.id,
          description: pillar.description,
          resources: [],
          questions: []
        }))
      ],
      {
        transaction
      }
    )

    const pillarsResources = await LensPillarResource.bulkCreate(
      pillars.flatMap((pillar, idx) => [
        ...(s.pillars[idx].resources?.map(resource => {
          return {
            pillarId: pillar.id,
            url: resource.url,
            description: resource.description
          }
        }) ?? [])
      ]),
      { transaction }
    )

    const questions = await LensPillarQuestion.bulkCreate(
      pillars.flatMap((pillar, idx) => [
        ...s.pillars[idx].questions.map(question => {
          return {
            pillarId: pillar.id,
            ref: question.id,
            name: question.title,
            description: question.description
          }
        })
      ]),
      { transaction }
    )

    const questionResources = await LensPillarQuestionResource.bulkCreate(
      pillars.flatMap((pillar, a) => [
        ...questions.flatMap((question, b) => [
          ...(s.pillars[a].questions[b]?.resources?.map(resource => {
            return {
              questionId: question.id,
              url: resource.url,
              description: resource.description
            }
          }) ?? [])
        ])
      ]),
      { transaction }
    )

    const choices = await LensPillarChoice.bulkCreate(
      pillars.flatMap((pillar, a) => [
        ...questions.flatMap((question, b) => [
          ...s.pillars[a].questions[b].choices.map(choice => {
            return {
              questionId: question.id,
              ref: choice.id,
              name: choice.title
            }
          })
        ])
      ]),
      { transaction }
    )

    return { ...lens.dataValues }
  })
}

export const findOneLensPillarQuestion = async (
  opts: z.infer<typeof LensesGetQuestionSchema>
) =>
  await LensPillarQuestion.findOne({
    include: [
      { model: LensPillarChoice },
      { model: LensPillarQuestionResource }
    ],
    where: { id: opts }
  })

export async function findAndCountLenses({
  offset = 0,
  limit = 10
}: Pagination) {
  const lenses = await Lens.findAndCountAll({
    order: [['name', 'DESC']],
    offset,
    limit
  })

  return lenses
}
