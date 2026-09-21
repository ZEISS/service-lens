import "server-only"

import { and, count, eq } from "drizzle-orm"

import { db } from "@/db"
import {
  assignEnvironmentSchema,
  assignLensSchema,
  removeEnvironmentSchema,
  removeLensSchema,
  workloadDeleteSchema,
  workloadEnvironment,
  workloadInsertSchema,
  workloadLens,
  workloads,
  assignProfileSchema,
  removeProfileSchema,
  workloadProfile,
} from "@/db/schema"
import type {
  TWorkloadAssignEnvironmentSchema,
  TWorkloadAssignLensSchema,
  TWorkloadDeleteSchema,
  TWorkloadInsertSchema,
  TWorkloadRemoveEnvironmentSchema,
  TWorkloadRemoveLensSchema,
  TWorkloadAssignProfileSchema,
  TWorkloadRemoveProfileSchema,
} from "@/db/schemas/workload"
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

export const assignEnvironment = async (input: TWorkloadAssignEnvironmentSchema) => {
  const parsed = await assignEnvironmentSchema.parseAsync(input)
  const result = await db.insert(workloadEnvironment).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const assignLens = async (input: TWorkloadAssignLensSchema) => {
  const parsed = await assignLensSchema.parseAsync(input)
  const result = await db.insert(workloadLens).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const assignProfile = async (input: TWorkloadAssignProfileSchema) => {
  const parsed = await assignProfileSchema.parseAsync(input)
  const result = await db.insert(workloadProfile).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const removeProfile = async (input: TWorkloadRemoveProfileSchema) => {
  const parsed = await removeProfileSchema.parseAsync(input)
  const result = await db
    .delete(workloadProfile)
    .where(
      and(
        eq(workloadProfile.workloadId, parsed.workloadId),
        eq(workloadProfile.profileId, parsed.profileId),
      ),
    )
    .returning()
  return takeFirstOrNull(result)
}

export const removeLens = async (input: TWorkloadRemoveLensSchema) => {
  const parsed = await removeLensSchema.parseAsync(input)
  const result = await db
    .delete(workloadLens)
    .where(and(eq(workloadLens.workloadId, parsed.workloadId), eq(workloadLens.lensId, parsed.lensId)))
    .returning()
  return takeFirstOrNull(result)
}

export const removeEnvironment = async (input: TWorkloadRemoveEnvironmentSchema) => {
  const parsed = await removeEnvironmentSchema.parseAsync(input)
  const result = await db
    .delete(workloadEnvironment)
    .where(
      and(
        eq(workloadEnvironment.workloadId, parsed.workloadId),
        eq(workloadEnvironment.environmentId, parsed.environmentId),
      ),
    )
    .returning()
  return takeFirstOrNull(result)
}

export const deleteWorkload = async (input: TWorkloadDeleteSchema) => {
  const parsed = await workloadDeleteSchema.parseAsync(input)
  await db.delete(workloads).where(eq(workloads.id, parsed.id))
}

export const getWorkloadById = async (id: string) =>
  await db.query.workloads.findFirst({
    where: { id },
    with: { environments: true, lenses: true, profiles: true, tags: true },
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
