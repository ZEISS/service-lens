import { z } from 'zod'

export const rhfActionDeleteLensSchema = z.string().uuid()
export type RHfActionDeleteLens = z.infer<typeof rhfActionDeleteLensSchema>

export const rhfActionPublishLensSchema = z.string().uuid()
export type RhfActionPublishLens = z.infer<typeof rhfActionPublishLensSchema>
