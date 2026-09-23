import "server-only"

import type { SQL } from "drizzle-orm"
import { and, count, eq, ilike } from "drizzle-orm"

import { db } from "@/db"
import {
  insertLensWithPillarsAndQuestionsSchema,
  lensDeleteSchema,
  lenses,
  lensPillarQuestionChoices,
  lensPillarQuestionResources,
  lensPillarQuestionRisks,
  lensPillarQuestions,
  lensPillars,
  type TLensDeleteSchema,
  type TLensWithPillarsSchema,
} from "@/db/schema"
import { takeFirstOrNull } from "@/db/utils"

import type { paginationParams } from "./pagination"

export type getLensesSchema = ReturnType<typeof paginationParams.parse>

export async function getLenses(input: getLensesSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const filters: SQL[] = []

      if (input.search) {
        filters.push(ilike(lenses.name, `${input.search}%`))
      }

      const data = await tx
        .select()
        .from(lenses)
        .where(and(...filters))
        .limit(input.perPage)
        .offset(offset)

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
    const lens = await db.query.lenses.findFirst({
      where: { id },
      with: {
        lensPillars: {
          with: {
            questions: {
              with: {
                choices: true,
                risks: true,
                resources: true,
              },
            },
          },
        },
      },
    })
    return lens
  } catch {
    return null
  }
}

export type TLensWithPillarsQuestions = Awaited<ReturnType<typeof getLensById>>

export const insertLensWithPillarsAndQuestions = async (input: TLensWithPillarsSchema) =>
  await db.transaction(async (tx) => {
    const parsed = await insertLensWithPillarsAndQuestionsSchema.parseAsync(input)
    const result = await tx.insert(lenses).values(parsed).returning()
    const lens = takeFirstOrNull(result)
    if (!lens) return null
    // TODO: Make this nice and collapse it into an easier step process.
    parsed.pillars?.forEach(async (pillar) => {
      const result = await tx
        .insert(lensPillars)
        .values({ ...pillar, lensId: lens?.id })
        .onConflictDoNothing()
        .returning()
      const lensPillar = takeFirstOrNull(result)
      if (!lensPillar) return

      pillar.questions?.forEach(async (question) => {
        const result = await tx
          .insert(lensPillarQuestions)
          .values({ ...question, pillarId: lensPillar.id })
          .onConflictDoNothing()
          .returning()

        question.choices?.forEach(async (choice) => {
          await tx
            .insert(lensPillarQuestionChoices)
            .values({ ...choice, questionId: result[0].id })
            .onConflictDoNothing()
            .returning()
        })

        question.resources?.forEach(async (resource) => {
          await tx
            .insert(lensPillarQuestionResources)
            .values({ ...resource, questionId: result[0].id })
            .onConflictDoNothing()
            .returning()
        })

        question.risks?.forEach(async (risk) => {
          await tx
            .insert(lensPillarQuestionRisks)
            .values({ ...risk, questionId: result[0].id })
            .onConflictDoNothing()
            .returning()
        })
      })
    })

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
