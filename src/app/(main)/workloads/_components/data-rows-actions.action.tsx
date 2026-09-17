"use server"

import { deleteWorkload } from "@/db/queries/workloads"
import { workloadDeleteSchema } from "@/db/schema"
import { revalidatePath } from "next/cache"
import "server-only"
import { z } from "zod"
import type { DeleteWorkloadSchema } from "./data-rows-actions.schema"

export async function deleteWorkloadAction(_: DeleteWorkloadSchema, data: FormData) {
  const values = {
    id: data.get("id") as string,
  }

  const result = workloadDeleteSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await deleteWorkload(result.data)
  } catch (_error) {
    return {
      success: false,
    }
  }

  revalidatePath("/workloads")

  return {
    success: true,
  }
}
