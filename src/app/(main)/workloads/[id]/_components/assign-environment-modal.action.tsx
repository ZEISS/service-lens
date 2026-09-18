"use server"

import { revalidatePath } from "next/cache"

import { insertWorkload, assignEnvironment } from "@/db/queries/workloads"
import { type TWorkload, assignEnvironmentSchema } from "@/db/schema"

import "server-only"

import { z } from "zod"

import type { AssignEnvironmentModalFormState } from "./assign-environment-modal.schema"

export async function assignEnvironmentAction(_: AssignEnvironmentModalFormState, data: FormData) {
  const values = {
    workloadId: data.get("workloadId") as string,
    environmentId: data.get("environmentId") as string,
  }

  const result = assignEnvironmentSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await assignEnvironment(result.data)
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
