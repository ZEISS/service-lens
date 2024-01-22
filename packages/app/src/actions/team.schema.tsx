import { z } from 'zod'

export const rhfActionDeleteTeamSchema = z.string().trim().min(3).max(128)
export type RHfActionDeleteTeam = z.infer<typeof rhfActionDeleteTeamSchema>
