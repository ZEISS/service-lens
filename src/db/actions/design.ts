"use server"

import { deleteDesign } from "@/db/queries/designs"
import { revalidatePath } from "next/cache"
import { z } from "zod"
import { type TDeleteDesignAction, deleteDesignSchema } from "./schema"

export const deleteAction: TDeleteDesignAction = async (_, data) => {
  try {
    const parsedData = deleteDesignSchema.parse(Object.fromEntries(data.entries()))
    await deleteDesign({ id: parsedData.id })
    revalidatePath("/designs", "page")
  } catch (err) {
    console.error(err)
    if (err instanceof z.ZodError) {
      return {
        success: false,
        errors: err.issues.map((issue) => ({ path: issue.path.join("."), message: issue.message })),
      }
    }
  }
  return {
    success: true,
    data: null,
  }
}
