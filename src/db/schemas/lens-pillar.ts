import { pgTable } from "@/db/utils"
import { timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { lenses } from "./lens"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

export const lensPillars = pgTable("lens_pillars", {
  id: uuid().primaryKey().defaultRandom(),
  ref: varchar({ length: 255 }).notNull(),
  name: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }).notNull(),
  lensId: uuid().notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export const insertLensPillarSchema = createInsertSchema(lensPillars, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
}).omit({
  createdAt: true,
  deletedAt: true,
  id: true,
  lensId: true,
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
