import { z } from 'zod'

export const FindAndCountTeamsSchema = z.object({
  limit: z.number().min(0).max(100).default(10),
  offset: z.number().min(0).default(0)
})
export const FindOneTeamSchema = z.string().uuid()
export const CreateTeamSchema = z.object({
  name: z.string().min(3).max(128),
  description: z.string().min(3).max(255).optional(),
  contactEmail: z.string().email().optional()
})
