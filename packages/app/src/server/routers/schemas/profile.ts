import { PaginationSchema } from './pagination'
import { z } from 'zod'

export const ProfileListSchema = PaginationSchema
export const ProfileGetSchema = z.string().uuid()
