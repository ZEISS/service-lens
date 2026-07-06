import type { TLensDeleteSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type DeleteLensSchema = ZodFormState<TLensDeleteSchema> | null
