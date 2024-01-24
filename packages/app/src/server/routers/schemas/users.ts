import { z } from 'zod'
import { PaginationSchema } from './pagination'

export const UserGetSchema = z.string().uuid()
export type UserGetSchema = z.infer<typeof UserGetSchema>

export const UserTeamsListSchema = z.string().uuid()
export type UserTeamsList = z.infer<typeof UserTeamsListSchema>

export const ListUsersSchema = PaginationSchema
export type ListUsers = z.infer<typeof ListUsersSchema>
