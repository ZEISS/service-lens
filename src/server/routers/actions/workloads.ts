import { protectedProcedure } from '../../trpc'
import {
  WorkloadGetSchema,
  WorkloadGetQuestionSchema,
  WorkloadListSchema,
  WorkloadDeleteSchema,
  WorkloadGetLensQuestionSchema,
  WorkloadGetAnswerSchema
} from '../schemas/workload'
import {
  getWorkload as gw,
  findWorkloadLensAnswer,
  findAndCountWorkloads,
  deleteWorkload as dt,
  getWorkloadLensQuestion,
  countWorkloads
} from '@/db/services/workloads'

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
