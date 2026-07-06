import type { TProfileInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AddProfileModalFormState = ZodFormState<TProfileInsertSchema> | null
