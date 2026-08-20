"use server"

import { insertLensWithPillarsAndQuestionsSchema, type TLens } from "@/db/schema"
import { insertLensWithPillarsAndQuestions } from "@/db/queries/lenses"
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

  const result = insertLensWithPillarsAndQuestionsSchema.safeParse({
    raw: spec.data,
    name: spec.data?.name,
    version: spec.data?.version,
    description: spec.data?.description,
    pillars: spec.data?.pillars,
  })

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    console.error("Validation errors:", errors.properties)

    return {
      values,
      errors,
      success: false,
    }
  }

  let lens: TLens | null = null

  try {
    lens = await insertLensWithPillarsAndQuestions(result.data)
  } catch (error) {
    console.error("insert error", error)

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
