import type { TNewDesign } from "@/db/schema"
import { z } from "zod"

export const newDesignSchema = z.object({
  title: z.string().min(1).max(255),
})

export type TNewDesignFormValues = z.infer<typeof newDesignSchema>
export const defaultValues: Partial<TNewDesign> = {
  title: "",
}
