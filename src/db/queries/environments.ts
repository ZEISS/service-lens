import "server-only"

import { count, eq } from "drizzle-orm"

import { db } from "@/db"
import {
  environmentDeleteSchema,
  environmentInsertSchema,
  environments,
  type TEnvironmentDeleteSchema,
  type TEnvironmentInsertSchema,
} from "@/db/schema"
import { takeFirstOrNull } from "@/db/utils"

import type { paginationParams } from "./pagination"

export type getEnvironmentsSchema = ReturnType<typeof paginationParams.parse>

export async function getEnvironments(input: getEnvironmentsSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const data = await tx.select().from(environments).limit(input.perPage).offset(offset)

      const total = await tx
        .select({
          count: count(),
        })
        .from(environments)
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

export const findEnvironmentById = async (id: string) => {
  const result = await db.select().from(environments).where(eq(environments.id, id))
  return takeFirstOrNull(result)
}

export const insertEnvironment = async (input: TEnvironmentInsertSchema) => {
  const parsed = await environmentInsertSchema.parseAsync(input)
  const result = await db.insert(environments).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const deleteEnvironment = async (input: TEnvironmentDeleteSchema) => {
  const parsed = await environmentDeleteSchema.parseAsync(input)
  await db.delete(environments).where(eq(environments.id, parsed.id))
}

export const getTotalNumberOfEnvironments = async () => {
  try {
    const result = await db
      .select({
        count: count(),
      })
      .from(environments)
      .execute()
      .then((res) => res[0]?.count ?? 0)

    return result
  } catch {
    return 0
  }
}
