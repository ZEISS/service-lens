import { pgTable } from "@/db/utils"
import { timestamp, uuid, varchar, pgEnum } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"
import { z } from "zod"

export const riskEnum = pgEnum("risk", ["NO_RISK", "LOW_RISK", "MEDIUM_RISK", "HIGH_RISK"]);

export const lensPillarQuestionRisks = pgTable("lens_pillars_questions_risks", {
  id: uuid().primaryKey().defaultRandom(),
  risk: riskEnum().default("NO_RISK"),
  condition: varchar({ length: 1024 }).notNull(),
  questionsId: uuid().notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export const insertLensPillarQuestionRisksSchema = createInsertSchema(lensPillarQuestionRisks, {
  risk: () => z.enum(["NO_RISK", "LOW_RISK", "MEDIUM_RISK", "HIGH_RISK"]),
  condition: (schema) => schema.min(1, "Condition is required").max(1024, "Condition must be at most 1024 characters"),
}).omit({
  createdAt: true,
  deletedAt: true,
  id: true,
  questionsId: true,
})

export type TLensPillarQuestionRisk = typeof lensPillarQuestionRisks.$inferSelect
export type TNewLensPillarQuestionRisk = typeof lensPillarQuestionRisks.$inferInsert

export const lensPillarQuestionRiskSelectSchema = createSelectSchema(lensPillarQuestionRisks)
export const lensPillarQuestionRiskDeleteSchema = createSelectSchema(lensPillarQuestionRisks).pick({
  id: true,
})

export type TLensPillarQuestionRiskInsertSchema = ReturnType<typeof lensPillarQuestionRiskSelectSchema.parse>
export type TLensPillarQuestionRiskSelectSchema = ReturnType<typeof lensPillarQuestionRiskSelectSchema.parse>
export type TLensPillarQuestionRiskDeleteSchema = ReturnType<typeof lensPillarQuestionRiskDeleteSchema.parse>
