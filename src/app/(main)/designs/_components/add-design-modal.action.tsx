"use server"

import { insertDesign } from "@/db/queries/designs"
import { designInsertSchema, type TDesign } from "@/db/schemas/design"
import { redirect } from "next/navigation"
import "server-only"
import { z } from "zod"
import type { AddDesignFormState } from "./add-design-modal.schema"

export async function createDesignAction(_: AddDesignFormState, data: FormData) {
  const values = {
    title: data.get("title") as string,
  }

  const result = designInsertSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  let design: TDesign

  try {
    design = await insertDesign(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  return redirect(`/designs/${design.id}`)
}
