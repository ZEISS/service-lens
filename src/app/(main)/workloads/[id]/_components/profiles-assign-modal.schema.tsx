import type { TWorkloadAssignProfileSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AssignProfilesModalFormState = ZodFormState<TWorkloadAssignProfileSchema> | null
