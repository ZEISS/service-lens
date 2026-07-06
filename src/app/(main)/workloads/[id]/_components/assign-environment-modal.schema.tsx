import type { TWorkloadInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AddWorkloadModalFormState = ZodFormState<TWorkloadInsertSchema> | null
