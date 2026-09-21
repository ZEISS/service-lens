import type { TWorkloadAssignLensSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AssignLensesModalFormState = ZodFormState<TWorkloadAssignLensSchema> | null
