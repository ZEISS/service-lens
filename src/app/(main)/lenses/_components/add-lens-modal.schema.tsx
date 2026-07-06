import type { TLensInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AddLensModalFormState = ZodFormState<TLensInsertSchema> | null
