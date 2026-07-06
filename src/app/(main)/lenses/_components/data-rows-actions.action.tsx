"use server"

import { deleteLens } from "@/db/queries/lenses"
import { lensDeleteSchema } from "@/db/schema"
import { revalidatePath } from "next/cache"
import "server-only"
import { z } from "zod"
import type { DeleteLensSchema } from "./data-rows-actions.schema"

export async function deleteLensAction(_: DeleteLensSchema, data: FormData) {
  const values = {
    id: data.get("id") as string,
  }

  const result = lensDeleteSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await deleteLens(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  revalidatePath("/lenses")

  return {
    success: true,
  }
}
