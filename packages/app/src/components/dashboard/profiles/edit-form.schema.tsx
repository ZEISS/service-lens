import { z } from 'zod'

export const rhfActionSchema = z.object({
  selectedChoices: z.record(z.string(), z.array(z.string()).min(1))
})

export type EditProfileFormValues = z.infer<typeof rhfActionSchema>
export const defaultValues: Partial<EditProfileFormValues> = {}
