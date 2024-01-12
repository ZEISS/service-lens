import { z } from 'zod'

export const FindOneUserSchema = z.string().uuid()
export type FindOneUserSchema = z.infer<typeof FindOneUserSchema>
