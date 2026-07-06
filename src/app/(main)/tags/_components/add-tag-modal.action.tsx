"use server"

import { insertTag } from "@/db/queries/tags"
import { tagInsertSchema, type TTag } from "@/db/schema"
import { redirect } from "next/navigation"
import "server-only"
import { z } from "zod"
import type { AddTagModalFormState } from "./add-tag-modal.schema"

export async function createTagAction(_: AddTagModalFormState, data: FormData) {
  const values = {
    name: data.get("name") as string,
    value: data.get("value") as string,
  }

  const result = tagInsertSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  let tag: TTag | null = null

  try {
    tag = await insertTag(result.data)
  } catch (error) {
    console.error("Error inserting tag:", error)

    return {
      success: false,
    }
  }

  if (!tag) {
    return {
      success: false,
    }
  }

  return redirect(`/tags/${tag?.id}`)
}
