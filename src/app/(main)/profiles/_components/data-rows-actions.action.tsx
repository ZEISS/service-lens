"use server"

import { deleteEnvironment } from "@/db/queries/environments"
import { environmentDeleteSchema } from "@/db/schema"
import { revalidatePath } from "next/cache"
import "server-only"
import { z } from "zod"
import type { DeleteEnvironmentSchema } from "./data-rows-actions.schema"

export async function deleteEnvironmentAction(_: DeleteEnvironmentSchema, data: FormData) {
  const values = {
    id: data.get("id") as string,
  }

  const result = environmentDeleteSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await deleteEnvironment(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  revalidatePath("/environments")

  return {
    success: true,
  }
}
