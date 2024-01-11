import { z } from 'zod'

export const rhfActionSchema = z.object({
  name: z.string().min(3).max(128),
  description: z.string().min(10).max(256).optional(),
  contactEmail: z.string().email().optional()
})

export type NewTeamFormValues = z.infer<typeof rhfActionSchema>
export const defaultValues: Partial<NewTeamFormValues> = {
  name: ''
}
