import "server-only"

import { count, eq } from "drizzle-orm"

import { db } from "@/db"
import { lensDeleteSchema, lenses, lensInsertSchema, type TLensDeleteSchema, type TLensInsertSchema } from "@/db/schema"
import { takeFirstOrNull } from "@/db/utils"

import type { paginationParams } from "./pagination"

export type getLensesSchema = ReturnType<typeof paginationParams.parse>

export async function getLenses(input: getLensesSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const data = await tx.select().from(lenses).limit(input.perPage).offset(offset)

      const total = await tx
        .select({
          count: count(),
        })
        .from(lenses)
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

export async function getLensById(id: string) {
  try {
    const lens = await db.select().from(lenses).where(eq(lenses.id, id)).execute().then(takeFirstOrNull)
    return lens
  } catch {
    return null
  }
}

export const insertLens = async (input: TLensInsertSchema) => {
  const parsed = await lensInsertSchema.parseAsync(input)
  const result = await db.insert(lenses).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const deleteLens = async (input: TLensDeleteSchema) => {
  const parsed = await lensDeleteSchema.parseAsync(input)
  await db.delete(lenses).where(eq(lenses.id, parsed.id))
}

export const getTotalNumberOfLenses = async () => {
  try {
    const result = await db
      .select({
        count: count(),
      })
      .from(lenses)
      .execute()
      .then((res) => res[0]?.count ?? 0)

    return result
  } catch {
    return 0
  }
}
