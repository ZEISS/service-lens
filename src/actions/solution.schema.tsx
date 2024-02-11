import { z } from 'zod'

export const rhfActionDeleteSolutionSchema = z.string().trim().uuid()
export type RHfActionDeleteSolution = z.infer<
  typeof rhfActionDeleteSolutionSchema
>
