import type { TWorkloadReviewInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type UpsertWorkloadReviewFormState = ZodFormState<TWorkloadReviewInsertSchema> | null
