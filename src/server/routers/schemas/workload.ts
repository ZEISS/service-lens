import { z } from 'zod'
import { PaginationSchema } from './pagination'
import { TeamsGetBySlugSchema } from '../schemas/teams'

export const WorkloadDeleteSchema = z.string().uuid()
export const WorkloadGetSchema = z.string().uuid()
export const WorkloadGetQuestionSchema = z.object({
  workloadId: z.string(),
  questionId: z.string()
})
export const WorkloadListSchema = PaginationSchema
export const WorkloadAddSchema = z.object({
  name: z.string().min(3).max(256),
  description: z.string().min(10).max(2048),
  spec: z.string()
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

export const ListWorkloadLensSchema = z
  .object({ id: z.string().trim().uuid() })
  .and(PaginationSchema)
export type ListWorkloadLens = z.infer<typeof ListWorkloadLensSchema>

export const ListWorkloadByTeamSlug = TeamsGetBySlugSchema.and(PaginationSchema)
export type ListWorkloadByTeamSlug = z.infer<typeof ListWorkloadByTeamSlug>
