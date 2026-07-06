"use server"

import { db } from "@/db"
import { designInsertSchema, designs } from "@/db/schema"
import "server-only"
import { z } from "zod"

export type CreateDesignActionState = {
  errors: z.ZodIssue[]
  success: boolean
  designId?: string
}

export async function createDesignAction(prev: CreateDesignActionState, state: FormData) {
  try {
    const newDesign = {
      title: state.get("title") as string,
    }

    const parsedDesign = designInsertSchema.parse(newDesign)
    const [item] = await db.insert(designs).values(parsedDesign).returning({ insertedId: designs.id })

    return { errors: [], success: true, designId: item.insertedId }
  } catch (err) {
    if (err instanceof z.ZodError) {
      return { errors: err.issues, success: false }
    }

    return { errors: [], success: false }
  }
}
