import { z } from "zod"

export interface ActionFormData<T extends Record<string, string | File>> extends FormData {
  get<K extends keyof T>(key: Extract<K, string>): T[K]
}

export type Result<T> =
  | null
  | {
      success: true
      data: T
    }
  | {
      success: false
      errors: Array<{ path: string; message: string }>
    }

export type Action<O> = (prev: Result<O>, data: FormData) => Promise<Result<O>>

export const deleteDesignSchema = z.object({
  id: z.uuid(),
})

export type TDeleteDesignSchema = z.infer<typeof deleteDesignSchema>
export type TDeleteDesignAction = Action<null>
