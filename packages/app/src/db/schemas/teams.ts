import { z } from 'zod'
import { PaginationSchema } from './pagination'

const reservedSlugs = ['app', 'admin', 'www', 'admin']

export const FindAndCountTeamsSchema = z.object({
  limit: z.number().min(0).max(100).default(10),
  offset: z.number().min(0).default(0)
})
export const FindOneTeamSchema = z.string().uuid()
export const CreateTeamSchema = z.object({
  name: z.string().min(3).max(128),
  slug: z
    .string()
    .min(3)
    .max(128)
    .refine(slug => !reservedSlugs.includes(slug), {
      message: "Slug can't be one of reserved slugs"
    }),
  userId: z.string().uuid(),
  description: z.string().min(3).max(255).optional(),
  contactEmail: z.string().email().optional()
})
export type CreateTeamSchema = z.infer<typeof CreateTeamSchema>

export const DestroyTeamSchema = z.string().uuid()
export type DestroyTeam = z.infer<typeof DestroyTeamSchema>

export const FindOneTeamByNameSlug = z.object({
  slug: z.string().trim().toLowerCase().min(3).max(128)
})
export type FindOneTeamByNameSlug = z.infer<typeof FindOneTeamByNameSlug>

export const FineOneTeamAndUsersByNameSlugSchema =
  FindOneTeamByNameSlug.and(PaginationSchema)
export type FineOneTeamAndUsersByNameSlug = z.infer<
  typeof FineOneTeamAndUsersByNameSlugSchema
>

export const ListWorkloadsByTeamSlug =
  FindOneTeamByNameSlug.and(PaginationSchema)
export type ListWorkloadsByTeamSlug = z.infer<typeof ListWorkloadsByTeamSlug>
