import { z } from 'zod'

export const FindAndCountEnvironmentsSchema = z.object({
  limit: z.number().min(0).max(100).default(10),
  offset: z.number().min(0).default(0)
})
