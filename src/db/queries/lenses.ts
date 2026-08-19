import "server-only"

import { count, eq } from "drizzle-orm"

import { db } from "@/db"
import { lensDeleteSchema, lensPillars, lenses, insertLensWithPillarsSchema, type TLensDeleteSchema, type TLensWithPillarsSchema } from "@/db/schema"
import { takeFirstOrNull } from "@/db/utils"

import type { paginationParams } from "./pagination"

export type getLensesSchema = ReturnType<typeof paginationParams.parse>

export async function getLenses(input: getLensesSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const total = await db.select({ count: count() }).from(lenses).execute().then((res) => res[0]?.count ?? 0)
    const data = await db.query.lenses.findMany({ with: { lensPillars: true }, limit: input.perPage, offset })
    const pageCount = Math.ceil(total / input.perPage)
    return { data, pageCount }
  } catch {
    return { data: [], pageCount: 0 }
  }
}

export async function getLensById(id: string) {
  try {
    const lens = await db.query.lenses.findFirst({ where: { id }, with: { lensPillars: { where: { lensId: id } } } })
    return lens
  } catch {
    return null
  }
}

export const insertLensWithPillars = async (input: TLensWithPillarsSchema) =>
  await db.transaction(async (tx) => {
    const parsed = await insertLensWithPillarsSchema.parseAsync(input)
    const result = await tx.insert(lenses).values(parsed).returning()
    const lens = takeFirstOrNull(result)
    if (!lens) return null
    parsed.pillars?.forEach(async (pillar) => await tx.insert(lensPillars).values({ ...pillar, lensId: lens?.id }).onConflictDoNothing())
    return lens
  })

export const deleteLens = async (input: TLensDeleteSchema) =>
  await db.transaction(async (tx) => {
    const parsed = await lensDeleteSchema.parseAsync(input)
    await tx.delete(lenses).where(eq(lenses.id, parsed.id))
  })

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
