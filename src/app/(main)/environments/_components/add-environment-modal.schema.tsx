import type { TEnvironmentInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AddEnvironmentModalFormState = ZodFormState<TEnvironmentInsertSchema> | null
