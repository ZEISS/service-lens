import { z } from 'zod'

export const FindOnePermissionSchema = z.object({
  userId: z.string().uuid(),
  teamId: z.string().uuid(),
  permission: z.string()
})
export type FindOnePermissionSchema = z.infer<typeof FindOnePermissionSchema>
