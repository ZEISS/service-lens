import "server-only"

import { and, count, eq } from "drizzle-orm"

import { db } from "@/db"
import {
  assignEnvironmentSchema,
  assignLensSchema,
  assignProfileSchema,
  removeEnvironmentSchema,
  removeLensSchema,
  removeProfileSchema,
  workloadDeleteSchema,
  workloadEnvironment,
  workloadInsertSchema,
  workloadLens,
  workloadProfile,
  workloadReview,
  workloadReviewDeleteSchema,
  workloadReviewInsertSchema,
  workloadReviewUpdateSchema,
  workloads,
} from "@/db/schema"
import type {
  TWorkloadAssignEnvironmentSchema,
  TWorkloadAssignLensSchema,
  TWorkloadAssignProfileSchema,
  TWorkloadDeleteSchema,
  TWorkloadInsertSchema,
  TWorkloadRemoveEnvironmentSchema,
  TWorkloadRemoveLensSchema,
  TWorkloadRemoveProfileSchema,
  TWorkloadReviewDeleteSchema,
  TWorkloadReviewInsertSchema,
  TWorkloadReviewSelectSchema,
  TWorkloadReviewUpdateSchema,
  TWorkloadSelectSchema,
} from "@/db/schemas/workload"
import { takeFirstOrNull } from "@/db/utils"

import type { paginationParams } from "./pagination"

export type getWorkloadsSchema = ReturnType<typeof paginationParams.parse>
export type listWorkloadReviewsSchema = ReturnType<typeof paginationParams.parse>

// List all workloads with pagination.
//
// @param input - The pagination parameters.
// @returns The workloads and total count.
export async function listWorkloadsWithPagination(input: getWorkloadsSchema) {
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

// Get a workload by its ID.
//
// @param id - The workload ID.
// @returns The workload.
export const getWorkload = async ({ id }: TWorkloadSelectByIDSchema) =>
  await db.query.workloads.findFirst({
    where: { id },
    with: {
      lenses: {
        with: {
          lensPillars: {
            with: {
              questions: {
                with: {
                  choices: true,
                },
              },
            },
          },
        },
      },
      environments: true,
      reviews: true,
      profiles: true,
      tags: true,
    },
  })
export type TWorkloadSelectByIDSchema = Pick<TWorkloadSelectSchema, "id">

// Get workload with the reviews for the lens.
export const getWorkloadWithReviewsForLens = async ({ id, lensId }: TWorkloadSelectWithReviewsAndLensSchema) => await db.query.workloads.findFirst({
  where: { id },
  with: {
    lenses: {
      where: {
        id: lensId
      },
      with: {
        lensPillars: true
      },
    },
    reviews: {
      where: { lensId },
    },
  },
})

export type TWorkloadSelectWithReviewsAndLensSchema = Pick<TWorkloadSelectSchema, "id"> & { lensId: string }

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
    .where(and(eq(workloadProfile.workloadId, parsed.workloadId), eq(workloadProfile.profileId, parsed.profileId)))
    .returning()
  return takeFirstOrNull(result)
}

export const removeLens = async (input: TWorkloadRemoveLensSchema) => {
  const parsed = await removeLensSchema.parseAsync(input)
  const result = await db.transaction((tx) => tx.delete(workloadLens)
    .where(and(eq(workloadLens.workloadId, parsed.workloadId), eq(workloadLens.lensId, parsed.lensId)))
    .returning())
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

// Get workload with its reviews.
//
// @param id - The workload ID.
// @returns The workload with its reviews.
export const getWorkloadWithReview = async ({ id }: TWorkloadWithReviewSelectSchema) =>
  await db.query.workloads.findFirst({
    where: { id },
    with: {
      reviews: true,
    },
  })
export type TWorkloadWithReviewSelectSchema = Pick<TWorkloadSelectSchema, "id">
export type TWorkloadWithReview = Awaited<ReturnType<typeof getWorkloadWithReview>>

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

export const deleteWorkloadReview = async (input: TWorkloadReviewDeleteSchema) => {
  const parsed = await workloadReviewDeleteSchema.parseAsync(input)
  return await db.transaction(async (tx) => {
    await tx.delete(workloadReview).where(eq(workloadReview.id, parsed.id))
  })
}

export type TWorkloadReviewSelectByLensId = Pick<TWorkloadReviewSelectSchema, "workloadId" | "lensId">
export const getWorkloadReviewByLensId = async ({ workloadId, lensId }: TWorkloadReviewSelectByLensId) =>
  await db.query.workloadReview.findFirst({
    where: { workloadId, lensId },
  })

export const insertWorkloadReview = async (input: TWorkloadReviewInsertSchema) => {
  const parsed = await workloadReviewInsertSchema.parseAsync(input)
  const result = await db.insert(workloadReview).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const updateWorkloadReview = async (input: TWorkloadReviewUpdateSchema) => await db.transaction(async (tx) => {
  const parsed = await workloadReviewUpdateSchema.parseAsync(input)
  return await tx.update(workloadReview).set(parsed).where(eq(workloadReview.id, parsed.id)).returning({id: workloadReview.id, notes: workloadReview.notes})
})
