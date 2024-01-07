import { Profile } from '@/db/models/profile'
import { WorkloadLensesAnswer } from '@/db/models/workload-lenses-answers'
import { LensPillarChoice } from '@/db/models/lens-pillar-choices'
import { Workload } from '@/db/models/workload'
import { WorkloadEnvironment } from '@/db/models/workload-environment'
import { Lens } from '@/db/models/lens'
import { WorkloadLens } from '@/db/models/workload-lens'
import { Environment } from '@/db/models/environment'
import { LensPillar } from '@/db/models/lens-pillars'
import { LensPillarQuestion } from '@/db/models/lens-pillar-questions'
import { WorkloadLensesAnswerChoice } from '@/db/models/workload-lenses-answers-choices'
import { v4 as uuidv4 } from 'uuid'
import { z } from 'zod'
import type { WorkloadCreationAttributes } from '../models/workload'
import sequelize from '../config/config'
import {
  WorkloadLensQuestionSchema,
  WorkloadGetLensAnswer,
  WorkloadLensAnswerAddSchema
} from '../schemas/workload'
import { Op } from 'sequelize'

export const findWorkloadLensAnswer = async (
  opts: z.infer<typeof WorkloadGetLensAnswer>
) =>
  await WorkloadLensesAnswer.findOne({
    where: { ...opts },
    include: [LensPillarChoice]
  })

export const countWorkloads = async () => await Workload.count()

export async function createWorkload({
  name,
  description,
  profilesId,
  environmentsIds,
  lensesIds
}: WorkloadCreationAttributes & {
  environmentsIds: bigint[]
  lensesIds: string[]
}) {
  return await sequelize.transaction(async transaction => {
    const id = uuidv4()

    const workload = await Workload.create(
      {
        id,
        profilesId,
        name,
        description
      },
      { transaction }
    )

    const items = Array.from(environmentsIds).map(id => ({
      environmentId: id,
      workloadId: workload.id
    }))
    await WorkloadEnvironment.bulkCreate(items, { transaction })

    await WorkloadLens.bulkCreate(
      Array.from(lensesIds).map(lensId => ({
        workloadId: workload.id,
        lensId
      })),
      { transaction }
    )

    return workload.dataValues
  })
}

export const deleteWorkload = async (id: string) =>
  await Workload.update({ deletedAt: new Date(Date.now()) }, { where: { id } })

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

export const addLensAnswer = async (
  opts: z.infer<typeof WorkloadLensAnswerAddSchema>
) =>
  await sequelize.transaction(async transaction => {
    const [answer] = await WorkloadLensesAnswer.upsert(
      {
        ...opts
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
      Array.from(opts.selectedChoices).map(id => ({
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

export const updateWorkloadAnswer = async ({
  answerId,
  doesNotApply,
  doesNotApplyReason
}: {
  answerId: string
  doesNotApply: boolean
  doesNotApplyReason: string
}) =>
  await sequelize.transaction(async transaction => {
    // WorkloadLensesAnswer.upsert
    // const answer = await WorkloadLensPillarAnswer.findOne({
    //   where: { id: answerId },
    //   transaction
    // })
    // if (!answer) {
    //   throw Error('Answer not found')
    // }
    // answer.doesNotApply = doesNotApply
    // answer.doesNotApplyReason = doesNotApplyReason
    // await answer?.save({ transaction })
  })

export const getWorkload = async (id: string) =>
  await Workload.findOne({
    where: { id },
    include: [Profile, Environment, Lens, WorkloadLens]
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
    include: [Profile, Environment],
    order: [['name', 'DESC']],
    offset,
    limit
  })

  return workloads
}
