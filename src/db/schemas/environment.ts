import { relations } from "drizzle-orm"
import { bigint, timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

import { pgTable } from "@/db/utils"

import { tags } from "./tag"
import { workloadEnvironment } from "./workload"

export const environments = pgTable("environment", {
  id: uuid().primaryKey().defaultRandom(),
  name: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TEnvironment = typeof environments.$inferSelect
export type TNewEnvironment = typeof environments.$inferInsert

export const environmentTag = pgTable("environment_tag", {
  environmentId: uuid()
    .notNull()
    .references(() => environments.id, { onDelete: "cascade" }),
  tagId: bigint({ mode: "bigint" })
    .notNull()
    .references(() => tags.id, { onDelete: "cascade" }),
})

export const environmentRelations = relations(environments, ({ many }) => ({
  workloads: many(workloadEnvironment),
  tags: many(tags),
}))

export const environmentInsertSchema = createInsertSchema(environments, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
  description: (schema) =>
    schema.min(1, "Description is required").max(1024, "Description must be at most 1024 characters"),
}).pick({
  name: true,
  description: true,
})

export const environmentSelectSchema = createSelectSchema(environments)
export const environmentDeleteSchema = createSelectSchema(environments).pick({
  id: true,
})

export type TEnvironmentInsertSchema = ReturnType<typeof environmentInsertSchema.parse>
export type TEnvironmentSelectSchema = ReturnType<typeof environmentSelectSchema.parse>
export type TEnvironmentDeleteSchema = ReturnType<typeof environmentDeleteSchema.parse>
