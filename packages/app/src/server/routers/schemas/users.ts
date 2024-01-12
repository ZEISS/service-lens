import { z } from 'zod'

export const UserGetSchema = z.string().uuid()
