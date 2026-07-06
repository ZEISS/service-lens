"use server"

import { insertLens } from "@/db/queries/lenses"
import { lensInsertSchema, type TLens } from "@/db/schema"
import { lensSpecSchema } from "@/lib/spec"
import { redirect } from "next/navigation"
import "server-only"
import * as z from "zod"
import type { AddLensModalFormState } from "./add-lens-modal.schema"

export async function createLensAction(_: AddLensModalFormState, data: FormData) {
  const values = {
    spec: data.get("spec") as File,
  }

  const buffer = await values.spec.arrayBuffer()
  const json = new TextDecoder().decode(buffer)
  const spec = lensSpecSchema.safeParse(JSON.parse(json))

  const result = lensInsertSchema.safeParse({
    raw: spec.data,
    name: spec.data?.name,
    version: spec.data?.version,
    description: spec.data?.description,
  })

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    console.error("Validation errors:", errors.properties?.version?.errors)

    return {
      values,
      errors,
      success: false,
    }
  }

  let lens: TLens | null = null

  try {
    lens = await insertLens(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  if (!lens) {
    return {
      success: false,
    }
  }

  return redirect(`/lenses/${lens?.id}`)
}
