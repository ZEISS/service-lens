import type { TWorkloadReviewInsertSchema, TWorkloadReviewUpdateSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type UpsertWorkloadReviewFormState = ZodFormState<TWorkloadReviewInsertSchema> | null
export type UpdateWorkloadReviewFormState = ZodFormState<TWorkloadReviewUpdateSchema> | null
