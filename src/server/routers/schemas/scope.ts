import { z } from 'zod'

export const ScopeSchema = z.object({
  team: z.string().trim().toLowerCase().min(3).max(128)
})
export type Scope = z.infer<typeof ScopeSchema>
