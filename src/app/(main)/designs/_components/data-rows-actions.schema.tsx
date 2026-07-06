import type { TDesignSelectSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type DeleteDesignSchema = ZodFormState<TDesignSelectSchema> | null
