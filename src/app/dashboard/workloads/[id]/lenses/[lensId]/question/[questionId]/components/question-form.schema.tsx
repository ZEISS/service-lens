import { z } from 'zod'

export const rhfActionSchema = z.object({
  // answerId: z.bigint().readonly(),
  workloadId: z.string().uuid().readonly(),
  lensPillarQuestionId: z.string().readonly(),
  selectedChoices: z.array(z.string()).min(1).default([]),
  doesNotApply: z.boolean().optional(),
  doesNotApplyReason: z.string().optional(),
  notes: z.string().min(10).max(2048).optional()
})
