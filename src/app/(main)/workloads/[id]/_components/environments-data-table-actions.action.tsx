"use server"

import { deleteProfile } from "@/db/queries/profiles"
import { profileDeleteSchema } from "@/db/schema"
import { revalidatePath } from "next/cache"
import "server-only"
import { z } from "zod"
import type { DeleteProfileSchema } from "./environments-data-table-actions.schema"

export async function deleteProfileAction(_: DeleteProfileSchema, data: FormData) {
  const values = {
    id: data.get("id") as string,
  }

  const result = profileDeleteSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await deleteProfile(result.data)
  } catch (_error) {
    return {
      success: false,
    }
  }

  revalidatePath("/profiles")

  return {
    success: true,
  }
}
