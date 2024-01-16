import { z } from 'zod'

export const generalFormSchema = z.object({
  organization: z.string().min(3).max(256).optional(),
  description: z.string().min(10).max(2048)
})

export type GeneralFormValues = z.infer<typeof generalFormSchema>
export const defaultValues: Partial<GeneralFormValues> = {}
