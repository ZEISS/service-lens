import { z } from 'zod'

export const rhfActionSchema = z.object({
  solutionId: z.string().uuid().readonly(),
  body: z.string().min(1).default('')
})
