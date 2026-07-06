import "server-only"

import { db } from "@/db"
import { profiles } from "@/db/schema"
import type { TProfile, TProfileDeleteSchema, TProfileInsertSchema } from "@/db/schemas/profile"
import { profileDeleteSchema, profileInsertSchema } from "@/db/schemas/profile"
import { takeFirstOrNull } from "@/db/utils"
import { count, eq } from "drizzle-orm"
import type { paginationParams } from "./pagination"

export type getProfilesSchema = ReturnType<typeof paginationParams.parse>

export async function getProfiles(input: getProfilesSchema) {
  try {
    const offset = (input.page - 1) * input.perPage
    const { data, total } = await db.transaction(async (tx) => {
      const data = await tx.select().from(profiles).limit(input.perPage).offset(offset)

      const total = await tx
        .select({
          count: count(),
        })
        .from(profiles)
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

export const insertProfile = async (input: TProfileInsertSchema) => {
  const parsed = await profileInsertSchema.parseAsync(input)
  const result = await db.insert(profiles).values(parsed).returning()
  return takeFirstOrNull(result)
}

export const deleteProfile = async (input: TProfileDeleteSchema) => {
  const parsed = await profileDeleteSchema.parseAsync(input)
  await db.delete(profiles).where(eq(profiles.id, parsed.id))
}

export const getProfileById = async (id: string): Promise<TProfile | null> => {
  try {
    const result = await db.select().from(profiles).where(eq(profiles.id, id)).limit(1)
    return result[0] || null
  } catch {
    return null
  }
}
