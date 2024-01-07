import { z } from 'zod'
import { PaginationSchema } from './pagination'

export const WorkloadDeleteSchema = z.string().uuid()
export const WorkloadGetSchema = z.string().uuid()
export const WorkloadGetQuestionSchema = z.object({
  workloadId: z.string(),
  questionId: z.string()
})
export const WorkloadListSchema = PaginationSchema
export const WorkloadAddSchema = z.object({
  name: z.string().min(3).max(256),
  description: z.string().min(10).max(2048)
})
export const WorkloadGetLensQuestionSchema = z.object({
  workloadId: z.string().uuid(),
  lensId: z.string().uuid(),
  questionId: z.string()
})
export const WorkloadGetAnswerSchema = z.object({
  workloadId: z.string().uuid(),
  lensPillarQuestionId: z.string()
})
