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

export const AddWorkloadLensAnswerSchema = z.discriminatedUnion(
  'doesNotApply',
  [
    z.object({
      workloadId: z.string().uuid().readonly(),
      lensPillarQuestionId: z.string().readonly(),
      selectedChoices: z.array(z.coerce.bigint()).optional(),
      doesNotApply: z.literal<boolean>(true),
      doesNotApplyReason: z.string(),
      notes: z.string().min(10).max(2048).optional()
    }),
    z.object({
      workloadId: z.string().uuid().readonly(),
      lensPillarQuestionId: z.string().readonly(),
      selectedChoices: z.array(z.coerce.bigint()).min(1).default([]),
      doesNotApply: z.literal<boolean>(false),
      doesNotApplyReason: z.string().optional(),
      notes: z.string().min(10).max(2048).optional()
    })
  ]
)

export type AddWorkloadLens = z.infer<typeof AddWorkloadLensAnswerSchema>

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
