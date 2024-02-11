import { z } from 'zod'

export const rhfActionNewSolutionSchema = z.object({
  title: z.string().min(1).max(256).default(''),
  description: z.string().optional(),
  body: z.string().min(1).default('')
})

export type NewSolutionFormValues = z.infer<typeof rhfActionNewSolutionSchema>
export const defaultValues: Partial<NewSolutionFormValues> = {
  title: '',
  description: '',
  body: ''
}
