import { protectedProcedure } from '../../trpc'
import {
  WorkloadGetSchema,
  WorkloadListSchema,
  WorkloadDeleteSchema,
  WorkloadGetLensQuestionSchema,
  WorkloadGetAnswerSchema,
  ListWorkloadByTeamSlug
} from '../schemas/workload'
import {
  getWorkload as gw,
  findWorkloadLensAnswer,
  findAndCountWorkloads,
  deleteWorkload as dt,
  getWorkloadLensQuestion,
  countWorkloads,
  listWorkloadByTeamSlug
} from '@/db/services/workloads'
import { router } from '@/server/trpc'

export const getWorkload = protectedProcedure
  .input(WorkloadGetSchema)
  .query(async opts => await gw(opts.input))

export const getWorkloadAnswer = protectedProcedure
  .input(WorkloadGetAnswerSchema)
  .query(async opts => findWorkloadLensAnswer(opts.input))

export const listWorkloads = protectedProcedure
  .input(WorkloadListSchema)
  .query(async opts => await findAndCountWorkloads({ ...opts.input }))

export const deleteWorkload = protectedProcedure
  .input(WorkloadDeleteSchema)
  .query(async opts => await dt(opts.input))

export const findWorkloadLensQuestion = protectedProcedure
  .input(WorkloadGetLensQuestionSchema)
  .query(async opts => await getWorkloadLensQuestion(opts.input))

export const totalWorkloads = protectedProcedure.query(
  async _ => await countWorkloads()
)

export const listByTeam = protectedProcedure
  .input(ListWorkloadByTeamSlug)
  .query(async opts => await listWorkloadByTeamSlug({ ...opts.input }))

export const workloadsRouter = router({
  get: getWorkload,
  list: listWorkloads,
  listByTeam: listByTeam,
  delete: deleteWorkload
})
