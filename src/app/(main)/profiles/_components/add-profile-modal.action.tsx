"use server"

import { insertProfile } from "@/db/queries/profiles"
import { environmentInsertSchema, type TProfile } from "@/db/schema"
import { redirect } from "next/navigation"
import "server-only"
import { z } from "zod"
import type { AddProfileModalFormState } from "./add-profile-modal.schema"

export async function createProfileAction(_: AddProfileModalFormState, data: FormData) {
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

  let profile: TProfile | null = null

  try {
    profile = await insertProfile(result.data)
  } catch (error) {
    return {
      success: false,
    }
  }

  if (!profile) {
    return {
      success: false,
    }
  }

  return redirect(`/profiles/${profile?.id}`)
}
