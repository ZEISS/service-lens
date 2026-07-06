"use server"

import { deleteDesign } from "@/db/queries/designs"
import { designDeleteSchema } from "@/db/schema"
import { revalidatePath } from "next/cache"
import "server-only"
import { z } from "zod"
import type { DeleteDesignSchema } from "./data-rows-actions.schema"

export async function deleteDesignAction(_: DeleteDesignSchema, data: FormData) {
  const values = {
    id: data.get("id") as string,
  }

  const result = designDeleteSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await deleteDesign(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  revalidatePath("/designs")

  return {
    success: true,
  }
}
