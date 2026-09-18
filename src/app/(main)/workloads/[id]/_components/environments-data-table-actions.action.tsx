"use server"

import { removeEnvironmentSchema } from "@/db/schema"
import { revalidatePath } from "next/cache"
import "server-only"
import { z } from "zod"
import type { RemoveEnvironmentSchema } from "./environments-data-table-actions.schema"
import { removeEnvironment } from "@/db/queries/workloads"

export async function removeEnvironmentAction(_: RemoveEnvironmentSchema, data: FormData) {
  const values = {
    workloadId: data.get("workloadId") as string,
    environmentId: data.get("environmentId") as string,
  }

  const result = removeEnvironmentSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await removeEnvironment(result.data)
  } catch (_error) {
    return {
      success: false,
    }
  }

  revalidatePath(`/workloads/${result.data.workloadId}`)

  return {
    success: true,
  }
}
