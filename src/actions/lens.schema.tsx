import { z } from 'zod'

export const rhfActionDeleteLensSchema = z.string().uuid()
export const rhfActionPublishLensSchema = z.string().uuid()
