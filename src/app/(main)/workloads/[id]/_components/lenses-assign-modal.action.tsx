"use server"

import { revalidatePath } from "next/cache"

import { assignLens } from "@/db/queries/workloads"
import { assignLensSchema } from "@/db/schema"

import "server-only"

import { z } from "zod"

import type { AssignLensesModalFormState } from "./lenses-assign-modal.schema"

export async function lensesAssignAction(_: AssignLensesModalFormState, data: FormData) {
  const values = {
    workloadId: data.get("workloadId") as string,
    lensId: data.get("lensId") as string,
  }

  const result = assignLensSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await assignLens(result.data)
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
