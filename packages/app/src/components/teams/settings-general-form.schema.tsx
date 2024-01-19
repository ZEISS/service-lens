import { z } from 'zod'

export const settingsGeneralFormSchema = z.object({
  name: z.string().min(3).max(256),
  description: z.string().min(10).max(2048).optional()
})

export type SettingGeneralFormValues = z.infer<typeof settingsGeneralFormSchema>
export const defaultValues: Partial<SettingGeneralFormValues> = {}
