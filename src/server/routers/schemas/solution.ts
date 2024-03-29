import { z } from 'zod'
import { PaginationSchema } from './pagination'
import { TeamsGetBySlugSchema } from '../schemas/teams'

export const SolutionListSchema = PaginationSchema
export const SolutionAddSchema = z.object({
  title: z.string().min(3).max(256),
  description: z.string().min(10).max(2048),
  body: z.string()
})
export const SolutionDeleteSchema = z.bigint()
export const SolutionGetSchema = z.string().uuid()
export const SolutionCommentDeleteSchema = z.bigint()
export const SolutionTemplateListSchema = PaginationSchema
export const SolutionTemplateGetSchema = z.string()
export const SolutionTemplateDeleteSchema = z.bigint()
export const SolutionMakeCopySchema = z.string().uuid()

export const DestroySolutionSchema = z.string().trim().uuid()
export type DestroySolution = z.infer<typeof DestroySolutionSchema>

export const ListSolutionByTeamSlug = TeamsGetBySlugSchema.and(PaginationSchema)
export type ListSolutionByTeamSlug = z.infer<typeof ListSolutionByTeamSlug>
