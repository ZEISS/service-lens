import { z } from 'zod'
import { FindOneTeamByNameSlug } from './teams'
import { PaginationSchema } from './pagination'

export const FindAndCountProfilesSchema = z.object({
  limit: z.number().min(0).max(100).default(10),
  offset: z.number().min(0).default(0)
})
export const FindAllProfilesQuestionsSchema = z.object({})
export const FindOneProfileSchema = z.string().uuid()
export const DestroyProfileSchema = z.string().uuid()
export const CreateProfileSchema = z.object({
  name: z.string().min(3).max(255),
  description: z.string().min(3).max(255),
  selectedChoices: z.record(z.string(), z.array(z.string()).min(1))
})
export const MakeCopyProfileSchema = z.string().uuid()

export const ListProfileByTeamSlug = FindOneTeamByNameSlug.and(PaginationSchema)
export type ListProfileByTeamSlug = z.infer<typeof ListProfileByTeamSlug>
