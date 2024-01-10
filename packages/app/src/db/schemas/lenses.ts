import { z } from 'zod'

export const LensesGetSchema = z.string().uuid()
export const LensesDeleteSchema = z.string().uuid()
export const LensesPublishSchema = z.string().uuid()
export const LensesGetQuestionSchema = z.string()
export const LensesAddSchema = z.object({
  name: z.string().min(1).max(256),
  description: z.string().min(10).max(2048),
  spec: z.string()
})
