import type { TProfileSelectSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type DeleteProfileSchema = ZodFormState<TProfileSelectSchema> | null
