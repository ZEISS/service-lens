"use server"

import { revalidatePath } from "next/cache"

import { deleteEnvironment } from "@/db/queries/environments"
import { tagDeleteSchema } from "@/db/schema"
import "server-only"

import { z } from "zod"

import { deleteTag } from "@/db/queries/tags"

import type { DeleteTagSchema } from "./data-rows-actions.schema"

export async function deleteTagAction(_: DeleteTagSchema, data: FormData) {
  const values = {
    id: data.get("id") as string,
  }

  const result = tagDeleteSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await deleteTag(result.data)
  } catch (_error) {
    return {
      success: false,
    }
  }

  revalidatePath("/environments")

  return {
    success: true,
  }
}
