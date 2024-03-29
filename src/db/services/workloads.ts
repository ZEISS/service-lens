import { Profile } from '@/db/models/profile'
import { WorkloadLensAnswer } from '@/db/models/workload-lenses-answers'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import { Workload } from '@/db/models/workload'
import { Lens } from '@/db/models/lens'
import { WorkloadLens } from '@/db/models/workload-lens'
import { Environment } from '@/db/models/environment'
import { LensPillar } from '@/db/models/lens-pillars'
import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { WorkloadLensesAnswerChoice } from '@/db/models/workload-lenses-answers-choices'
import { z } from 'zod'
import { Tag } from '@/db/models/tags'
import { createContext, evalInScope } from '@/lib/eval'
import { Team } from '@/db/models/teams'
import sequelize from '../config/config'
import {
  WorkloadLensQuestionSchema,
  WorkloadGetLensAnswer,
  AddWorkloadLens,
  DestroyWorkload
} from '../schemas/workload'
import { Op } from 'sequelize'
import { LensPillarQuestionRisk } from '../models/lens-pillar-risks'
import type { ListWorkloadsByTeamSlug } from '../schemas/workload'
import type { WorkloadCreate } from '../schemas/workload'
import { Ownership } from '../models/ownership'
import { type ListWorkloadLens } from '@/db/schemas/workload'

export const findWorkloadLensAnswer = async (
  opts: z.infer<typeof WorkloadGetLensAnswer>
) =>
  await WorkloadLensAnswer.findOne({
    where: { ...opts },
    include: [LensPillarChoice]
  })

export const countWorkloads = async () => await Workload.count()

export const destroyWorkload = async (opts: DestroyWorkload) =>
  sequelize.transaction(
    async transaction =>
      await Workload.destroy({ where: { id: opts }, transaction })
  )

export const getWorkloadLensQuestion = async (
  opts: z.infer<typeof WorkloadLensQuestionSchema>
) =>
  await Workload.findOne({
    where: { id: opts.workloadId },
    include: [
      {
        model: Lens,
        include: [
          {
            model: LensPillar,
            include: [
              {
                model: LensPillarQuestion,
                where: { id: '1' }
              }
            ]
          }
        ]
      }
    ]
  })

export const addLensAnswer = async (opts: AddWorkloadLens) =>
  await sequelize.transaction(async transaction => {
    const question = await LensPillarQuestion.findOne({
      where: { id: opts.lensPillarQuestionId },
      include: [LensPillarQuestionRisk, LensPillarChoice],
      transaction
    })

    const ctx = {
      ...createContext(question?.questionAnswers),
      ...question?.questionAnswers
        ?.filter(answer => opts.selectedChoices?.includes(answer.id))
        ?.reduce((answers, answer) => ({ ...answers, [answer.ref]: true }), {})
    }

    const defaultCondition = question?.risks?.find(
      risk => risk.condition === 'default'
    )

    const risk =
      question?.risks?.reduce((prev, curr) => {
        if (curr.condition === 'default') return curr.risk

        try {
          const truthy = evalInScope(curr.condition, ctx)
          return truthy ? curr.risk ?? 'UNANSWERED' : prev
        } catch (error) {
          console.error(error)
          return prev
        }
      }, defaultCondition?.risk) ?? 'UNANSWERED'

    const [answer] = await WorkloadLensAnswer.upsert(
      {
        ...opts,
        risk: risk
      },
      { transaction }
    )

    await WorkloadLensesAnswerChoice.destroy({
      where: {
        [Op.and]: [
          {
            choiceId: { [Op.notIn]: opts.selectedChoices }
          },
          {
            answerId: answer.id
          }
        ]
      },
      transaction
    })

    await WorkloadLensesAnswerChoice.bulkCreate(
      Array.from(opts.selectedChoices ?? []).map(id => ({
        answerId: answer.id,
        choiceId: id
      })),
      {
        transaction,
        updateOnDuplicate: ['answerId', 'choiceId', 'deletedAt', 'updatedAt']
      }
    )

    return answer.dataValues
  })

export const getWorkload = async (id: string) =>
  await Workload.findOne({
    where: { id },
    include: [
      Profile,
      Environment,
      Lens,
      { model: WorkloadLensAnswer, include: [LensPillarChoice] }
    ]
  })

export type Pagination = {
  offset?: number
  limit?: number
}

export async function findAndCountWorkloads({
  offset = 0,
  limit = 10
}: Pagination) {
  const workloads = await Workload.findAndCountAll({
    include: [Profile, Environment, Tag],
    order: [['name', 'DESC']],
    offset,
    limit
  })

  return workloads
}

export const listWorkloadByTeamSlug = async (opts: ListWorkloadsByTeamSlug) =>
  await Workload.findAndCountAll({
    offset: opts.offset,
    limit: opts.limit,
    include: [{ model: Team, where: { slug: opts.slug } }, Profile]
  })

export const listWorkloadLens = async (opts: ListWorkloadLens) =>
  await Lens.findAndCountAll({
    offset: opts.offset,
    limit: opts.limit,
    include: [{ model: Workload, where: { id: opts.id }, include: [Team] }]
  })

export const createWorkload = async (opts: WorkloadCreate) =>
  await sequelize.transaction(async transaction => {
    const workload = await Workload.create(
      {
        name: opts.name,
        description: opts.description,
        profilesId: opts.profile
      },
      { transaction }
    )

    await WorkloadLens.bulkCreate(
      opts.lenses.map(lens => ({ workloadId: workload.id, lensId: lens })),
      { transaction }
    )

    const _ = await Ownership.create(
      {
        ownerId: opts.ownerId,
        resourceId: workload.id,
        resourceType: 'workload'
      },
      { transaction }
    )

    return workload
  })
