import { z } from 'zod'

export const rhfActionProfileSchema = z.object({
  name: z
    .string()
    .min(2, {
      message: 'Name must be at least 2 characters.'
    })
    .max(255, {
      message: 'Name must not be longer than 255 characters.'
    }),
  email: z
    .string({
      required_error: 'Please select an email to display.'
    })
    .email()
})

export type ProfileFormValues = z.infer<typeof rhfActionProfileSchema>
export const defaultValues: Partial<ProfileFormValues> = {}
