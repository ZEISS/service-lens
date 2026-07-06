import "server-only"

import { count, eq } from "drizzle-orm"

import { db } from "@/db"
import { designs } from "@/db/schema"
import {
  designDeleteSchema,
  designInsertSchema,
  type TDesign,
  type TDesignDeleteSchema,
  type TDesignInsertSchema,
} from "@/db/schemas/design"

import type { paginationParams } from "./pagination"

export type GetDesignsSchema = ReturnType<typeof paginationParams.parse>

export async function getDesigns(input: GetDesignsSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const data = await tx.select().from(designs).limit(input.perPage).offset(offset)

      const total = await tx
        .select({
          count: count(),
        })
        .from(designs)
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

export const insertDesign = async (input: TDesignInsertSchema) => {
  const parsed = await designInsertSchema.parseAsync(input)
  const result = await db.insert(designs).values(parsed).returning()
  return result[0]
}

export const deleteDesign = async (input: TDesignDeleteSchema) => {
  const parsed = await designDeleteSchema.parseAsync(input)
  await db.delete(designs).where(eq(designs.id, parsed.id))
}

export const getTotalNumberOfDesigns = async () => {
  try {
    const result = await db
      .select({
        count: count(),
      })
      .from(designs)
      .execute()
      .then((res) => res[0]?.count ?? 0)

    return result
  } catch {
    return 0
  }
}

export const updateDesign = async (id: string, input: Partial<TDesignInsertSchema>): Promise<TDesign | null> => {
  try {
    const updateData = designInsertSchema.partial().parse(input)
    const result = await db
      .update(designs)
      .set({ ...updateData, updatedAt: new Date() })
      .where(eq(designs.id, id))
      .returning()
    return result[0] || null
  } catch {
    return null
  }
}

export const getDesignById = async (id: string): Promise<TDesign | null> => {
  try {
    const result = await db.select().from(designs).where(eq(designs.id, id)).limit(1)
    return result[0] || null
  } catch {
    return null
  }
}
