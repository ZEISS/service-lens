"use server"

import { revalidatePath } from "next/cache"

import { assignProfile } from "@/db/queries/workloads"
import { assignProfileSchema } from "@/db/schema"

import "server-only"

import { z } from "zod"

import type { AssignProfilesModalFormState } from "./profiles-assign-modal.schema"

export async function profilesAssignAction(_: AssignProfilesModalFormState, data: FormData) {
  const values = {
    workloadId: data.get("workloadId") as string,
    profileId: data.get("profileId") as string,
  }

  const result = assignProfileSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await assignProfile(result.data)
  } catch (_error) {
    return {
      success: false,
    }
  }

  revalidatePath(`/workloads/${values.workloadId}`)

  return {
    success: true,
  }
}
