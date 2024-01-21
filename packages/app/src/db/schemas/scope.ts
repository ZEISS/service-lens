import { z } from 'zod'

export const ScopeResourceTypeSchema = z.union([
  z.literal('workload'),
  z.literal('solution'),
  z.literal('lens'),
  z.literal('profile')
])
export type ScopeResourceType = z.infer<typeof ScopeResourceTypeSchema>

export const ScopeOwnerIdSchema = z.string().uuid()
export type ScopeOwnerId = z.infer<typeof ScopeOwnerIdSchema>

export const ScopeResourceIdSchema = z.string().uuid()
export type ScopeResourceId = z.infer<typeof ScopeResourceIdSchema>

export const ScopeSchema = z.object({
  ownerId: ScopeOwnerIdSchema,
  resourceType: ScopeResourceTypeSchema,
  resourceId: ScopeResourceIdSchema.optional()
})
export type Scope = z.infer<typeof ScopeSchema>
