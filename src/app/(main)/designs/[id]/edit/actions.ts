"use server"

import { revalidatePath } from "next/cache"

import { getDesignById, updateDesign } from "@/db/queries/designs"
import "server-only"

import { z } from "zod"

const updateDesignSchema = z.object({
  title: z.string().min(5, "Title must be at least 5 characters").max(255, "Title must be at most 255 characters"),
  description: z.string().max(1024, "Description must be at most 1024 characters").optional().nullable(),
  body: z.string().optional().nullable(),
})

export type UpdateDesignFormData = z.infer<typeof updateDesignSchema>

export type UpdateDesignResult = {
  success: boolean
  error?: string
  data?: any
}

export async function getDesignByIdAction(id: string): Promise<{ success: boolean; data?: any; error?: string }> {
  try {
    const design = await getDesignById(id)

    if (!design) {
      return {
        success: false,
        error: "Design not found",
      }
    }

    return {
      success: true,
      data: design,
    }
  } catch (error) {
    console.error("Error fetching design:", error)
    return {
      success: false,
      error: "Failed to load design",
    }
  }
}

export async function updateDesignAction(id: string, data: UpdateDesignFormData): Promise<UpdateDesignResult> {
  try {
    // Validate the input data
    const validatedData = updateDesignSchema.parse(data)

    // Check if design exists
    const existingDesign = await getDesignById(id)
    if (!existingDesign) {
      return {
        success: false,
        error: "Design not found",
      }
    }

    // Update the design
    const updatedDesign = await updateDesign(id, {
      title: validatedData.title,
      description: validatedData.description,
      body: validatedData.body,
    })

    if (!updatedDesign) {
      return {
        success: false,
        error: "Failed to update design",
      }
    }

    // Revalidate the design pages
    revalidatePath("/designs")
    revalidatePath(`/designs/${id}`)
    revalidatePath(`/designs/${id}/edit`)

    return {
      success: true,
      data: updatedDesign,
    }
  } catch (error) {
    console.error("Error updating design:", error)

    if (error instanceof z.ZodError) {
      return {
        success: false,
        error: error.issues.map((e) => e.message).join(", "),
      }
    }

    return {
      success: false,
      error: "Failed to update design",
    }
  }
}
