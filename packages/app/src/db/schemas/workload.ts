import { z } from 'zod'
import { FindOneTeamByNameSlug } from './teams'
import { PaginationSchema } from './pagination'
import { rhfActionSetScopeSchema } from '@/actions/teams-switcher.schema'
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
    name: z.string(),
    description: z.string()
  })
  .and(ScopeSchema)
export type WorkloadCreate = z.infer<typeof WorkloadCreateSchema>
