import "server-only"

import { db } from "@/db"
import {
  tagDeleteSchema,
  tagInsertSchema,
  tags,
  type TTag,
  type TTagDeleteSchema,
  type TTagInsertSchema,
} from "@/db/schema"
import { takeFirstOrNull } from "@/db/utils"
import { count, eq } from "drizzle-orm"
import type { paginationParams } from "./pagination"

export type getTagsSchema = ReturnType<typeof paginationParams.parse>

export async function getTags(input: getTagsSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const data = await tx.select().from(tags).limit(input.perPage).offset(offset)

      const total = await tx
        .select({
          count: count(),
        })
        .from(tags)
        .execute()
        .then((res) => res[0]?.count ?? 0)

      return {
        data,
        total,
      }
    })

    const pageCount = Math.ceil(total / input.perPage)
    return { data, pageCount }
  } catch {
    return { data: [], pageCount: 0 }
  }
}

export const insertTag = async (input: TTagInsertSchema) => {
  const parsed = await tagInsertSchema.parseAsync(input)
  const result = await db.insert(tags).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const deleteTag = async (input: TTagDeleteSchema) => {
  const parsed = await tagDeleteSchema.parseAsync(input)
  await db.delete(tags).where(eq(tags.id, parsed.id))
}

export const getTagById = async (id: number): Promise<TTag | null> => {
  try {
    const result = await db.select().from(tags).where(eq(tags.id, id)).limit(1)
    return result[0] || null
  } catch {
    return null
  }
}
