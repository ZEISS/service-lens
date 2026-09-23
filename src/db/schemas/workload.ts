import { bigint, boolean, timestamp, uniqueIndex, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

import { pgTable } from "@/db/utils"

import { environments } from "./environment"
import { lenses } from "./lens"
import { profiles } from "./profile"
import { tags } from "./tag"

export const workloads = pgTable("workload", {
  id: uuid().primaryKey().defaultRandom(),
  name: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }).notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TWorkload = typeof workloads.$inferSelect
export type TNewWorkload = typeof workloads.$inferInsert

export const workloadLens = pgTable(
  "workload_lens",
  {
    workloadId: uuid()
      .notNull()
      .references(() => workloads.id, { onDelete: "cascade" }),
    lensId: uuid()
      .notNull()
      .references(() => lenses.id),
  },
  (table) => [uniqueIndex("uniqueWorkloadLens").on(table.workloadId, table.lensId)],
)

export const workloadProfile = pgTable(
  "workload_profile",
  {
    workloadId: uuid()
      .notNull()
      .references(() => workloads.id, { onDelete: "cascade" }),
    profileId: uuid()
      .notNull()
      .references(() => profiles.id),
  },
  (table) => [uniqueIndex("uniqueWorkloadProfile").on(table.workloadId, table.profileId)],
)

export const workloadEnvironment = pgTable(
  "workload_environment",
  {
    workloadId: uuid()
      .notNull()
      .references(() => workloads.id, { onDelete: "cascade" }),
    environmentId: uuid()
      .notNull()
      .references(() => environments.id),
  },
  (table) => [uniqueIndex("uniqueWorkloadEnvironment").on(table.workloadId, table.environmentId)],
)

export const workloadTag = pgTable(
  "workload_tag",
  {
    workloadId: uuid()
      .notNull()
      .references(() => workloads.id, { onDelete: "cascade" }),
    tagId: bigint({ mode: "bigint" })
      .notNull()
      .references(() => tags.id, { onDelete: "cascade" }),
  },
  (table) => [uniqueIndex("uniqueWorkloadTag").on(table.workloadId, table.tagId)],
)

// Answers to questions for a workload
export const workloadAnswer = pgTable("workload_answers", {
  id: uuid().primaryKey().defaultRandom(),
  workloadId: uuid().notNull(),
  lensId: uuid().notNull(),
  questionId: uuid().notNull(),
  notApplicable: boolean().notNull(),
  reason: varchar({ length: 1024 }),
  notes: varchar({ length: 1024 }),
})

// Choices for answer options for a workload
export const workloadAnswerChoices = pgTable(
  "workload_answer_choices",
  {
    answerId: uuid().notNull(),
    choice: uuid().notNull(),
  },
  (table) => [uniqueIndex("uniqueWorkloadAnswerChoice").on(table.answerId, table.choice)],
)

export const workloadInsertSchema = createInsertSchema(workloads, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
  description: (schema) =>
    schema.min(1, "Description is required").max(1024, "Description must be at most 1024 characters"),
}).pick({
  name: true,
  description: true,
})

export const assignLensSchema = createInsertSchema(workloadLens).pick({
  workloadId: true,
  lensId: true,
})

export const removeLensSchema = createSelectSchema(workloadLens).pick({
  workloadId: true,
  lensId: true,
})

export const assignEnvironmentSchema = createInsertSchema(workloadEnvironment).pick({
  workloadId: true,
  environmentId: true,
})

export const removeEnvironmentSchema = createSelectSchema(workloadEnvironment).pick({
  workloadId: true,
  environmentId: true,
})

export const workloadSelectSchema = createSelectSchema(workloads)
export const workloadDeleteSchema = createSelectSchema(workloads).pick({
  id: true,
})

export const assignProfileSchema = createInsertSchema(workloadProfile).pick({
  workloadId: true,
  profileId: true,
})

export const removeProfileSchema = createSelectSchema(workloadProfile).pick({
  workloadId: true,
  profileId: true,
})

export type TWorkloadInsertSchema = ReturnType<typeof workloadInsertSchema.parse>
export type TWorkloadSelectSchema = ReturnType<typeof workloadSelectSchema.parse>
export type TWorkloadDeleteSchema = ReturnType<typeof workloadDeleteSchema.parse>
export type TWorkloadAssignEnvironmentSchema = ReturnType<typeof assignEnvironmentSchema.parse>
export type TWorkloadRemoveEnvironmentSchema = ReturnType<typeof removeEnvironmentSchema.parse>
export type TWorkloadAssignLensSchema = ReturnType<typeof assignLensSchema.parse>
export type TWorkloadRemoveLensSchema = ReturnType<typeof removeLensSchema.parse>
export type TWorkloadAssignProfileSchema = ReturnType<typeof assignProfileSchema.parse>
export type TWorkloadRemoveProfileSchema = ReturnType<typeof removeProfileSchema.parse>
