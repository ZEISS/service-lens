"use server"

import { redirect } from "next/navigation"

import { insertWorkload } from "@/db/queries/workloads"
import { type TWorkload, workloadInsertSchema } from "@/db/schema"
import "server-only"

import { z } from "zod"

import type { AddWorkloadModalFormState } from "../../_components/add-workload-modal.schema"

export async function createWorkloadAction(_: AddWorkloadModalFormState, data: FormData) {
  const values = {
    name: data.get("name") as string,
    description: data.get("description") as string,
  }

  const result = workloadInsertSchema.safeParse(values)

  if (!result.success) {
    const errors = z.treeifyError(result.error)

    return {
      values,
      errors,
      success: false,
    }
  }

  let workload: TWorkload | null = null

  try {
    workload = await insertWorkload(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  if (!workload) {
    return {
      success: false,
    }
  }

  return redirect(`/workloads/${workload?.id}`)
}
