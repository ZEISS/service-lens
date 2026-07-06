import type { TTagInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AddTagModalFormState = ZodFormState<TTagInsertSchema> | null
