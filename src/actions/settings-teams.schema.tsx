import { z } from 'zod'

export const rhfActionDeleteTeamSchema = z.string().uuid()
export type RhfActionDelete = z.infer<typeof rhfActionDeleteTeamSchema>
