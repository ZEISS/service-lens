import { z } from 'zod'

export const PaginationSchema = z.object({
  limit: z.coerce.number().min(0).max(100).default(10),
  offset: z.coerce.number().min(0).default(0)
})

export type PaginationSchema = z.infer<typeof PaginationSchema>
