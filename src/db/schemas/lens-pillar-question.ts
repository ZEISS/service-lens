import { pgTable } from "@/db/utils"
import { timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { lensPillars } from "./lens-pillar"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

export const lensPillarQuestions = pgTable("lens_pillars_questions", {
  id: uuid().primaryKey().defaultRandom(),
  ref: varchar({ length: 255 }).notNull(),
  title: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }).notNull(),
  pillarId: uuid().notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export const insertLensPillarQuestionSchema = createInsertSchema(lensPillarQuestions, {
  title: (schema) => schema.min(1, "Title is required").max(255, "Title must be at most 255 characters"),
  ref: (schema) => schema.min(1, "Reference is required").max(255, "Reference must be at most 255 characters"),
  description: (schema) => schema.min(1, "Description is required").max(1024, "Description must be at most 1024 characters"),
}).omit({
  createdAt: true,
  deletedAt: true,
  id: true,
  pillarId: true,
})

export type TLensPillarQuestion = typeof lensPillarQuestions.$inferSelect
export type TNewLensPillarQuestion = typeof lensPillarQuestions.$inferInsert

export const lensPillarQuestionSelectSchema = createSelectSchema(lensPillarQuestions)
export const lensPillarQuestionDeleteSchema = createSelectSchema(lensPillarQuestions).pick({
  id: true,
})

export type TLensPillarQuestionInsertSchema = ReturnType<typeof lensPillarQuestionSelectSchema.parse>
export type TLensPillarQuestionSelectSchema = ReturnType<typeof lensPillarQuestionSelectSchema.parse>
export type TLensPillarQuestionDeleteSchema = ReturnType<typeof lensPillarQuestionDeleteSchema.parse>
