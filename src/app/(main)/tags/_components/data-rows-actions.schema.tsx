import type { TTagDeleteSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type DeleteTagSchema = ZodFormState<TTagDeleteSchema> | null
