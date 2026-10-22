"use server"

import { revalidatePath } from "next/cache"
import { headers } from "next/headers"

import { z } from "zod"

import { auth } from "@/lib/auth"

import { insertWorkloadReview } from "../queries/workloads"
import { workloadReviewInsertSchema } from "../schema"
import type { UpsertWorkloadReviewFormState } from "./workload.schema"

export async function insertWorkloadReviewAction(_: UpsertWorkloadReviewFormState,data: FormData) {
  const session = await auth.api.getSession({
    headers: await headers(),
  })

  if (!session?.user) throw new Error("Unauthorized")

  const values = {
    lensId: data.get("lensId") as string,
    workloadId: data.get("workloadId") as string,
    notes: data.get("notes") as string | undefined,
  }

  const result = workloadReviewInsertSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    console.log(errors.properties)

    return {
      values,
      errors,
      success: false,
    }
  }

  try {
    await insertWorkloadReview(result.data)
  } catch (error) {
    console.error(error)
    return {
      success: false,
    }
  }

  revalidatePath("/workloads")

  return {
    success: true,
  }
}
