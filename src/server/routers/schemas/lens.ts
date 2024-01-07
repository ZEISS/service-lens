import { z } from 'zod'
import { PaginationSchema } from './pagination'

export const LensDeleteSchema = z.string().uuid()
export const LensGetSchema = z.string().uuid()
export const LensGetQuestionSchema = z.string()
export const LensListSchema = PaginationSchema
