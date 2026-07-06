import type { TDesignInsertSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AddDesignFormState = ZodFormState<TDesignInsertSchema> | null
