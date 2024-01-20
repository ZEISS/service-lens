import { z } from 'zod'

export const rhfActionSetScopeSchema = z.string()
export type rhfActionSetScopeSchema = z.infer<typeof rhfActionSetScopeSchema>
