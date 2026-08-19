import { pgTable } from "@/db/utils"
import { timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { lensPillars } from "./lens-pillar"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

export const lensPillarQuestionResources = pgTable("lens_pillars_questions_resources", {
  id: uuid().primaryKey().defaultRandom(),
  ref: varchar({ length: 255 }).notNull(),
  url: varchar({ length: 1024 }).notNull(),
  description: varchar({ length: 1024 }).notNull(),
  questionsId: uuid().notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export const insertLensPillarQuestionRessourcesSchema = createInsertSchema(lensPillarQuestionResources, {
  url: (schema) => schema.url(),
  description: (schema) => schema.min(1, "Description is required").max(1024, "Description must be at most 1024 characters"),
}).omit({
  createdAt: true,
  deletedAt: true,
  id: true,
  questionsId: true,
})

export type TLensPillarQuestionResource = typeof lensPillarQuestionResources.$inferSelect
export type TNewLensPillarQuestionResource = typeof lensPillarQuestionResources.$inferInsert

export const lensPillarQuestionResourceSelectSchema = createSelectSchema(lensPillarQuestionResources)
export const lensPillarQuestionResourceDeleteSchema = createSelectSchema(lensPillarQuestionResources).pick({
  id: true,
})

export type TLensPillarQuestionResourceInsertSchema = ReturnType<typeof lensPillarQuestionResourceSelectSchema.parse>
export type TLensPillarQuestionResourceSelectSchema = ReturnType<typeof lensPillarQuestionResourceSelectSchema.parse>
export type TLensPillarQuestionResourceDeleteSchema = ReturnType<typeof lensPillarQuestionResourceDeleteSchema.parse>
