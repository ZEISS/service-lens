import { PaginationSchema } from './pagination'
import { z } from 'zod'

export const TeamsListSchema = PaginationSchema
export const TeamsGetSchema = z.string().uuid()
export const TeamsCreateSchema = z.object({
  name: z.string().min(3).max(128),
  description: z.string().min(10).max(256).optional(),
  contactEmail: z.string().email().optional()
})
