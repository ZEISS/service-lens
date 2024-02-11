import { z } from 'zod'

export const rhfActionDeleteUserSchema = z.string().uuid()
export type RhfActionDeleteUser = z.infer<typeof rhfActionDeleteUserSchema>
