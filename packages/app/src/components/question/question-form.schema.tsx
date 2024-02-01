import { z } from 'zod'

export const rhfActionSchema = z.discriminatedUnion('doesNotApply', [
  z.object({
    workloadId: z.string().uuid().readonly(),
    lensPillarQuestionId: z.string().readonly(),
    selectedChoices: z.array(z.coerce.bigint()).optional(),
    doesNotApply: z.literal<boolean>(true),
    doesNotApplyReason: z.string(),
    notes: z.string().min(10).max(2048).optional()
  }),
  z.object({
    workloadId: z.string().uuid().readonly(),
    lensPillarQuestionId: z.string().readonly(),
    selectedChoices: z.array(z.coerce.bigint()).min(1).default([]),
    doesNotApply: z.literal<boolean>(false),
    doesNotApplyReason: z.string().optional(),
    notes: z.string().min(10).max(2048).optional()
  })
])
