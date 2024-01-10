import { z } from 'zod'

export const rhfActionSchema = z.object({
  name: z.string().min(3, {}).default(''),
  lensesIds: z.array(z.string().uuid()).min(1).default([]),
  environmentsIds: z.array(z.bigint()).min(1).default([]),
  description: z
    .string()
    .min(10, {
      message: 'Description must be at least 30 characters.'
    })
    .max(2024, {
      message: 'Description must be less than 2024 characters.'
    })
    .default(''),
  profilesId: z.string().uuid().default(''),
  tags: z.array(z.string()).default([])
})
