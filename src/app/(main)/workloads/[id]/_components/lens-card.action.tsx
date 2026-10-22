"use server"

import { revalidatePath } from "next/cache"
import { headers } from "next/headers"

import { removeLens } from "@/db/queries/workloads"
import { removeLensSchema } from "@/db/schema"
import { auth } from "@/lib/auth"

export async function removeLensAction(data: FormData) {
  const session = await auth.api.getSession({
    headers: await headers(),
  })

  if (!session?.user) throw new Error("Unauthorized")

  const values = {
    workloadId: data.get("workloadId") as string,
    lensId: data.get("lensId") as string,
  }

  const result = removeLensSchema.safeParse(values)

  if (!result.success) {
    throw new Error(result.error.message)
  }

  await removeLens(result.data)

  revalidatePath("/workloads")
}
