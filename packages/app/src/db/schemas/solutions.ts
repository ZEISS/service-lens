import { z } from 'zod'
import { FindOneTeamByNameSlug } from './teams'
import { PaginationSchema } from './pagination'
import { ScopeSchema } from './scope'

export const FindAndCountSolutionsSchema = z.object({
  limit: z.number().min(0).max(100).default(10),
  offset: z.number().min(0).default(0)
})

export const SolutionsGetSchema = z.string().uuid()
export const SolutionCommentAddSchema = z.object({
  solutionId: z.string().uuid().readonly(),
  userId: z.string(),
  body: z.string().min(1)
})
export const SolutionCommentDeleteSchema = z.bigint()
export const FindAndCountSolutionTemplates = PaginationSchema
export const FindOneSolutionTemplate = z.string()
export const DestroySolutionSchema = z.string()
export const DestroySolutionTemplateSchema = z.bigint()
export const MakeCopySolutionTemplateSchema = z.bigint()
export const MakeCopySolutionSchema = z.string().uuid()

export const ListSolutionByTeamSlug =
  FindOneTeamByNameSlug.and(PaginationSchema)
export type ListSolutionByTeamSlug = z.infer<typeof ListSolutionByTeamSlug>

export const SolutionCreateSchema = z
  .object({
    title: z.string().trim(),
    body: z.string(),
    description: z.string().optional(),
    userId: z.string().uuid()
  })
  .and(ScopeSchema)
export type SolutionCreate = z.infer<typeof SolutionCreateSchema>
