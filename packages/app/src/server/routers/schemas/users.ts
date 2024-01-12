import { z } from 'zod'

export const UserGetSchema = z.string().uuid()
export type UserGetSchema = z.infer<typeof UserGetSchema>

export const UserTeamsListSchema = z.string().uuid()
export type UserTeamsListSchema = z.infer<typeof UserTeamsListSchema>
