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
  LensesGetQuestionSchema,
  LensesAddSchema
} from '../schemas/lenses'
import sequelize from '@/db/config/config'
import { z } from 'zod'
import { LensPillarResource } from '../models/lens-pillar-resources'
import { LensPillarQuestionRisk } from '../models/lens-pillar-risks'

export type Pagination = {
  offset?: number
  limit?: number
}

export const getLens = async (opts: z.infer<typeof LensesGetSchema>) =>
  await Lens.findOne({
    where: { id: opts },
    include: [
      {
        model: LensPillar,
        include: [
          {
            model: LensPillarQuestion,
            include: [LensPillarQuestionRisk, LensPillarChoice]
          }
        ]
      }
    ]
  })

export const deleteLens = async (opts: z.infer<typeof LensesDeleteSchema>) =>
  await Lens.destroy({
    where: { id: opts }
  })

export const publishLens = async (opts: z.infer<typeof LensesPublishSchema>) =>
  await Lens.update({ isDraft: false }, { where: { id: opts } })

export const createLens = async (opts: z.infer<typeof LensesAddSchema>) =>
  await sequelize.transaction(async transaction => {
    const spec = await Spec.parseAsync(JSON.parse(opts.spec))

    const lens = await Lens.create(
      {
        ...opts,
        version: spec.version,
        spec: spec,
        isDraft: true
      },
      {
        transaction
      }
    )

    const pillars = await LensPillar.bulkCreate(
      [
        ...spec.pillars.map(pillar => ({
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
        ...(spec.pillars[idx].resources?.map(resource => {
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
        ...spec.pillars[idx].questions.map(question => {
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

    const questionRisks = await LensPillarQuestionRisk.bulkCreate(
      pillars.flatMap((pillar, a) => [
        ...questions.flatMap((question, b) => [
          ...(spec.pillars[a].questions[b]?.risks?.map(risk => {
            return {
              questionId: question.id,
              risk: risk.risk,
              condition: risk.condition
            }
          }) ?? [])
        ])
      ]),
      { transaction }
    )

    const questionResources = await LensPillarQuestionResource.bulkCreate(
      pillars.flatMap((pillar, a) => [
        ...questions.flatMap((question, b) => [
          ...(spec.pillars[a].questions[b]?.resources?.map(resource => {
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
          ...spec.pillars[a].questions[b].choices.map(choice => {
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
