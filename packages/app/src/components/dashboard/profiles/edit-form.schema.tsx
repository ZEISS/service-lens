import { z } from 'zod'

export const rhfEditProfileActionSchema = z.record(
  z.string(),
  z.array(z.string()).min(1)
)
export type EditProfileFormValues = z.infer<typeof rhfEditProfileActionSchema>
export const defaultValues: Partial<EditProfileFormValues> = {}
