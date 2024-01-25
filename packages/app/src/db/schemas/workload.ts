import { z } from 'zod'
import { FindOneTeamByNameSlug } from './teams'
import { PaginationSchema } from './pagination'
import { ScopeSchema } from './scope'

export const WorkloadLensQuestionSchema = z.object({
  workloadId: z.string().uuid(),
  lensId: z.string().uuid(),
  questionId: z.string()
})

export const WorkloadGetLensAnswer = z.object({
  workloadId: z.string().uuid(),
  lensPillarQuestionId: z.string()
})

export const WorkloadLensAnswerAddSchema = z.object({
  workloadId: z.string().uuid(),
  lensPillarQuestionId: z.string(),
  selectedChoices: z.array(z.string()).default([]),
  doesNotApply: z.boolean().optional(),
  doesNotApplyReason: z.string().optional(),
  notes: z.string().optional()
})

export const ListWorkloadsByTeamSlug =
  FindOneTeamByNameSlug.and(PaginationSchema)
export type ListWorkloadsByTeamSlug = z.infer<typeof ListWorkloadsByTeamSlug>

export const WorkloadCreateSchema = z
  .object({
    name: z.string().trim(),
    description: z.string(),
    profile: z.string().trim().uuid(),
    lenses: z.array(z.string().trim().uuid())
  })
  .and(ScopeSchema)
export type WorkloadCreate = z.infer<typeof WorkloadCreateSchema>

export const ListWorkloadLensSchema = z
  .object({ id: z.string().trim().uuid() })
  .and(PaginationSchema)
export type ListWorkloadLens = z.infer<typeof ListWorkloadLensSchema>

export const DestroyWorkloadSchema = z.string().trim().uuid()
export type DestroyWorkload = z.infer<typeof DestroyWorkloadSchema>
