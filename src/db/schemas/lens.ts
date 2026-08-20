import { pgTable } from "@/db/utils"
import { integer, json, timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"
import { insertLensPillarSchema } from "./lens-pillar"
import { insertLensPillarQuestionSchema } from "./lens-pillar-question"
import { z } from "zod"

export const lenses = pgTable("lens", {
  id: uuid().primaryKey().defaultRandom(),
  name: varchar({ length: 255 }).notNull(),
  version: integer().notNull(),
  description: varchar({ length: 1024 }),
  raw: json("raw").notNull().$type<Record<string, any>>(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TLens = typeof lenses.$inferSelect
export type TNewLens = typeof lenses.$inferInsert

export const lensInsertSchema = createInsertSchema(lenses, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
  version: (schema) => schema.default(1),
  description: (schema) => schema.max(1024, "Description must be at most 1024 characters").optional(),
}).pick({
  name: true,
  version: true,
  raw: true,
})

export const insertLensWithPillarsSchema = lensInsertSchema.extend({
  pillars: z.array(insertLensPillarSchema).optional(),
})

export const insertLensWithPillarsAndQuestionsSchema = insertLensWithPillarsSchema.extend({
  questions: z.array(insertLensPillarQuestionSchema).optional(),
})

export const lensSelectSchema = createSelectSchema(lenses)
export const lensDeleteSchema = createSelectSchema(lenses).pick({
  id: true,
})

export type TLensInsertSchema = ReturnType<typeof lensInsertSchema.parse>
export type TLensSelectSchema = ReturnType<typeof lensSelectSchema.parse>
export type TLensDeleteSchema = ReturnType<typeof lensDeleteSchema.parse>
export type TLensWithPillarsSchema = ReturnType<typeof insertLensWithPillarsSchema.parse>
