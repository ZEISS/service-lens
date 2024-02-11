import { z } from 'zod'
import { FindOneTeamByNameSlug } from './teams'
import { PaginationSchema } from './pagination'
import { ScopeSchema } from './scope'

export const LensesGetSchema = z.string().uuid()

export const LensesDeleteSchema = z
  .object({
    lensId: z.string().trim().uuid()
  })
  .and(ScopeSchema)
export type LensDelete = z.infer<typeof LensesDeleteSchema>

export const LensesPublishSchema = z
  .object({
    lensId: z.string().trim().uuid()
  })
  .and(ScopeSchema)
export type LensPublish = z.infer<typeof LensesPublishSchema>

export const LensesGetQuestionSchema = z.string()

export const LensesAddSchema = z
  .object({
    name: z.string().min(1).max(256),
    description: z.string().min(10).max(2048),
    spec: z.string()
  })
  .and(ScopeSchema)
export type CreateLens = z.infer<typeof LensesAddSchema>

export const ListLensByTeamSlug = FindOneTeamByNameSlug.and(PaginationSchema)
export type ListLensByTeamSlug = z.infer<typeof ListLensByTeamSlug>
