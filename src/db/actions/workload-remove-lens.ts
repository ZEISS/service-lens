import { revalidatePath } from "next/cache"
import { headers } from "next/headers"

import { z } from "zod"

import { auth } from "@/lib/auth"
import type { ZodFormState } from "@/types"

import { removeLens } from "../queries/workloads"
import { removeLensSchema, type TWorkloadRemoveLensSchema } from "../schema"

export type TWorkloadRemoveLensschema = ZodFormState<TWorkloadRemoveLensSchema> | null
export async function removeLensAction(_: TWorkloadRemoveLensschema, data: FormData) {
  const session = await auth.api.getSession({
    headers: await headers(),
  })

  if (!session?.user) throw new Error("Unauthorized")

  const values = {
    id: data.get("id") as string,
    notes: data.get("notes") as string | undefined,
  }

  const result = removeLensSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await removeLens(result.data)
  } catch {
    return {
      success: false,
    }
  }

  revalidatePath("/workloads")

  return {
    success: true,
  }
}
