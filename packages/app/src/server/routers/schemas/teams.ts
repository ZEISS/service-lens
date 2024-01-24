import { PaginationSchema } from './pagination'
import { z } from 'zod'

export const TeamsListSchema = PaginationSchema
export type TeamsList = z.infer<typeof TeamsListSchema>

export const TeamsGetSchema = z.string().uuid()
export const TeamsCreateSchema = z.object({
  name: z.string().min(3).max(128),
  slug: z.string().trim().toLowerCase().min(3).max(128).default(''),
  description: z.string().min(10).max(256).optional(),
  contactEmail: z.string().email().optional()
})

export const TeamsDestroySlugSchema = z.string()
export type TeamsDestroySlug = z.infer<typeof TeamsDestroySlugSchema>

export const TeamsGetBySlugSchema = z.object({
  slug: z.string().trim().toLowerCase().min(3).max(128).default('')
})
export type TeamsGetBySlugSchema = z.infer<typeof TeamsGetBySlugSchema>

export const ListWorkloadByTeamSlug = TeamsGetBySlugSchema.and(PaginationSchema)
export type ListWorkloadByTeamSlug = z.infer<typeof ListWorkloadByTeamSlug>

export const GetTeamAndUsersByTeamSlug =
  TeamsGetBySlugSchema.and(PaginationSchema)
export type GetTeamAndUsersByTeamSlug = z.infer<
  typeof GetTeamAndUsersByTeamSlug
>
