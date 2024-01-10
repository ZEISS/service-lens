import { z } from 'zod'

export const rhfActionSchema = z.object({
  title: z.string().min(1).max(256).default(''),
  description: z.string().optional(),
  body: z.string().min(1).default('')
})
