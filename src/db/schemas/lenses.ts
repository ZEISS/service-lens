import { z } from 'zod'

export const LensesGetSchema = z.string().uuid()
export const LensesDeleteSchema = z.string().uuid()
export const LensesPublishSchema = z.string().uuid()
export const LensesGetQuestionSchema = z.string()
