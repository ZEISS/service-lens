import type { TWorkloadAssignEnvironmentSchema } from "@/db/schema"
import type { ZodFormState } from "@/types"

export type AssignEnvironmentModalFormState = ZodFormState<TWorkloadAssignEnvironmentSchema> | null
