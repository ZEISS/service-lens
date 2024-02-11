import { z } from 'zod'

const reservedSlugs = ['app', 'admin', 'www', 'admin']

export const settingsGeneralFormSchema = z.object({
  name: z.string().min(3).max(256),
  slug: z
    .string()
    .min(3)
    .max(128)
    .refine(slug => !reservedSlugs.includes(slug), {
      message: "Slug can't be one of reserved slugs."
    }),
  description: z.string().min(10).max(2048).optional()
})

export type SettingGeneralFormValues = z.infer<typeof settingsGeneralFormSchema>
export const defaultValues: Partial<SettingGeneralFormValues> = {}
