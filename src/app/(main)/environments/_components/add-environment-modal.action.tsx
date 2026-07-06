"use server"

import { insertEnvironment } from "@/db/queries/environments"
import { environmentInsertSchema, type TEnvironment } from "@/db/schema"
import { redirect } from "next/navigation"
import "server-only"
import { z } from "zod"
import type { AddEnvironmentModalFormState } from "./add-environment-modal.schema"

export async function createEnvironmentAction(_: AddEnvironmentModalFormState, data: FormData) {
  const values = {
    name: data.get("name") as string,
  }

  const result = environmentInsertSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  let environment: TEnvironment | null = null

  try {
    environment = await insertEnvironment(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  if (!environment) {
    return {
      success: false,
    }
  }

  return redirect(`/environments/${environment?.id}`)
}
