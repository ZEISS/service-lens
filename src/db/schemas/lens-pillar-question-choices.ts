import { pgTable } from "@/db/utils"
import { timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { lensPillars } from "./lens-pillar"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

export const lensPillarQuestionChoices = pgTable("lens_pillars_questions_choices", {
  id: uuid().primaryKey().defaultRandom(),
  ref: varchar({ length: 255 }).notNull(),
  title: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }).notNull(),
  questionId: uuid().notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export const insertLensPillarQuestionChoicesSchema = createInsertSchema(lensPillarQuestionChoices, {
  title: (schema) => schema.min(1, "Title is required").max(255, "Title must be at most 255 characters"),
  ref: (schema) => schema.min(1, "Reference is required").max(255, "Reference must be at most 255 characters"),
  description: (schema) =>
    schema.min(1, "Description is required").max(1024, "Description must be at most 1024 characters"),
}).omit({
  createdAt: true,
  deletedAt: true,
  id: true,
  questionId: true,
})

export type TLensPillarQuestionChoice = typeof lensPillarQuestionChoices.$inferSelect
export type TNewLensPillarQuestionChoice = typeof lensPillarQuestionChoices.$inferInsert

export const lensPillarQuestionChoiceSelectSchema = createSelectSchema(lensPillarQuestionChoices)
export const lensPillarQuestionChoiceDeleteSchema = createSelectSchema(lensPillarQuestionChoices).pick({
  id: true,
})

export type TLensPillarQuestionChoiceInsertSchema = ReturnType<typeof lensPillarQuestionChoiceSelectSchema.parse>
export type TLensPillarQuestionChoiceSelectSchema = ReturnType<typeof lensPillarQuestionChoiceSelectSchema.parse>
export type TLensPillarQuestionChoiceDeleteSchema = ReturnType<typeof lensPillarQuestionChoiceDeleteSchema.parse>
