import { z } from 'zod'
import { TeamsCreateSchema } from '@/server/routers/schemas/teams'

export type NewTeamFormValues = z.infer<typeof TeamsCreateSchema>
export const defaultValues: Partial<NewTeamFormValues> = {
  name: '',
  slug: ''
}
