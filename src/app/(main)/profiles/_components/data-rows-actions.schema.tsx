import type { TEnvironmentSelectSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type DeleteEnvironmentSchema = ZodFormState<TEnvironmentSelectSchema> | null
