import "server-only"

import { count, eq } from "drizzle-orm"

import { db } from "@/db"
import { workloadDeleteSchema, workloadInsertSchema, workloads } from "@/db/schema"
import type { TWorkloadDeleteSchema, TWorkloadInsertSchema } from "@/db/schemas/workload"
import { takeFirstOrNull } from "@/db/utils"

import type { paginationParams } from "./pagination"

export type getWorkloadsSchema = ReturnType<typeof paginationParams.parse>

export async function getWorkloads(input: getWorkloadsSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const data = await tx.select().from(workloads).limit(input.perPage).offset(offset)

      const total = await tx
        .select({
          count: count(),
        })
        .from(workloads)
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

export const insertWorkload = async (input: TWorkloadInsertSchema) => {
  const parsed = await workloadInsertSchema.parseAsync(input)
  const result = await db.insert(workloads).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const deleteWorkload = async (input: TWorkloadDeleteSchema) => {
  const parsed = await workloadDeleteSchema.parseAsync(input)
  await db.delete(workloads).where(eq(workloads.id, parsed.id))
}

export const getWorkloadById = async (id: string) => await db.query.workloads.findFirst({
    with: {
      environments: {
          with: {environment: true},
      },
    },
    where: eq(workloads.id, id),
  })

export const getTotalNumberOfWorkloads = async () => {
  try {
    const result = await db
      .select({
        count: count(),
      })
      .from(workloads)
      .execute()
      .then((res) => res[0]?.count ?? 0)
    return result
  } catch (e) {
    console.error("Error fetching total number of workloads:", e)
    return 0
  }
}
