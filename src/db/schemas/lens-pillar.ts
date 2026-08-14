import { description } from "@/app/(main)/dashboard/default/_components/chart-area-interactive"
import { pgTable } from "@/db/utils"
import { integer, json, timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"
import { de } from "zod/v4/locales"

export const lensPillars = pgTable("lens_pillars", {
  id: uuid().primaryKey().defaultRandom(),
  ref: varchar({ length: 255 }).notNull(),
  name: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }).notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TLensPillar = typeof lensPillars.$inferSelect
export type TNewLensPillar = typeof lensPillars.$inferInsert

export const lensPillarSelectSchema = createSelectSchema(lensPillars)
export const lensPillarDeleteSchema = createSelectSchema(lensPillars).pick({
  id: true,
})

export type TPillarInsertSchema = ReturnType<typeof lensPillarSelectSchema.parse>
export type TPillarSelectSchema = ReturnType<typeof lensPillarSelectSchema.parse>
export type TPillarDeleteSchema = ReturnType<typeof lensPillarSelectSchema.parse>
