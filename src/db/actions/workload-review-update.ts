"use server"

import { revalidatePath } from "next/cache"
import { headers } from "next/headers"

import { z } from "zod"

import { auth } from "@/lib/auth"

import { updateWorkloadReview } from "../queries/workloads"
import { workloadReviewUpdateSchema } from "../schema"
import type { UpdateWorkloadReviewFormState } from "./workload.schema"

export async function updateWorkloadReviewAction(_: UpdateWorkloadReviewFormState,data: FormData) {
  const session = await auth.api.getSession({
    headers: await headers(),
  })

  if (!session?.user) throw new Error("Unauthorized")

  const values = {
    id: data.get("id") as string,
    notes: data.get("notes") as string | undefined,
  }

  const result = workloadReviewUpdateSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await updateWorkloadReview(result.data)
  } catch  {
    return {
      success: false,
    }
  }

  revalidatePath("/workloads")

  return {
    success: true,
  }
}
