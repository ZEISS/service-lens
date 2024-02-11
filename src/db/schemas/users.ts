import { z } from 'zod'
import { PaginationSchema } from './pagination'

export const FindOneUserSchema = z.string().uuid()
export type FindOneUserSchema = z.infer<typeof FindOneUserSchema>

export const FindAndCountUsersSchema = PaginationSchema
export type FindAndCountUsers = z.infer<typeof FindAndCountUsersSchema>

export const DestroyUserSchema = z.string().uuid()
export type DestroyUser = z.infer<typeof DestroyUserSchema>
