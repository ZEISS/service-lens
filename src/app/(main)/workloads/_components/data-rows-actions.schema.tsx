import type { TWorkloadDeleteSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type DeleteWorkloadSchema = ZodFormState<TWorkloadDeleteSchema> | null
