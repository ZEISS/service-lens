import { PaginationSchema } from './pagination'
import { TeamsGetBySlugSchema } from '../schemas/teams'
import { z } from 'zod'

export const ProfileListSchema = PaginationSchema
export const ProfileGetSchema = z.string().uuid()

export const ListProfileByTeamSlug = TeamsGetBySlugSchema.and(PaginationSchema)
export type ListProfileByTeamSlug = z.infer<typeof ListProfileByTeamSlug>
