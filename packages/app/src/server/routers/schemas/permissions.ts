import { z } from 'zod'

export const PermissionGetSchema = z.object({
  userId: z.string().uuid(),
  teamId: z.string().uuid(),
  permission: z.string()
})
export type PermissionGetSchema = z.infer<typeof PermissionGetSchema>
