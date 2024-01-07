import { PaginationSchema } from '@/server/routers/schemas/pagination'
import { z } from 'zod'

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
