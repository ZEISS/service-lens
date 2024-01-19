import { z } from 'zod'
import { PaginationSchema } from './pagination'
import { TeamsGetBySlugSchema } from '../schemas/teams'

export const LensDeleteSchema = z.string().uuid()
export const LensGetSchema = z.string().uuid()
export const LensGetQuestionSchema = z.string()
export const LensListSchema = PaginationSchema

export const ListLensByTeamSlug = TeamsGetBySlugSchema.and(PaginationSchema)
export type ListLensByTeamSlug = z.infer<typeof ListLensByTeamSlug>
