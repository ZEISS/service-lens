import { Pagination } from './pagination'
import { TeamsGetBySlugSchema } from '../schemas/teams'
import { z } from 'zod'

export const ProfileListSchema = Pagination
export const ProfileGetSchema = z.string().uuid()

export const ListProfileByTeamSlug = TeamsGetBySlugSchema.and(Pagination)
export type ListProfileByTeamSlug = z.infer<typeof ListProfileByTeamSlug>
