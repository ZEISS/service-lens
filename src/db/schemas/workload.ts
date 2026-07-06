import { pgTable } from "@/db/utils"
import { relations } from "drizzle-orm"
import { bigint, timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"
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

export const workloadLens = pgTable("workload_lens", {
  workloadId: uuid()
    .notNull()
    .references(() => workloads.id, { onDelete: "cascade" }),
  lensId: uuid()
    .notNull()
    .references(() => lenses.id),
  createdAt: timestamp("created_at").defaultNow().notNull(),
  updatedAt: timestamp("updated_at").defaultNow().notNull(),
})

export const workloadProfile = pgTable("workload_profile", {
  workloadId: uuid()
    .notNull()
    .references(() => workloads.id, { onDelete: "cascade" }),
  profileId: uuid()
    .notNull()
    .references(() => profiles.id),
})

export const workloadEnvironment = pgTable("workload_environment", {
  id: bigint({ mode: "bigint" }).primaryKey(),
  workloadId: uuid()
    .notNull()
    .references(() => workloads.id, { onDelete: "cascade" }),
  environmentId: uuid()
    .notNull()
    .references(() => environments.id),
  createdAt: timestamp("created_at").defaultNow().notNull(),
  updatedAt: timestamp("updated_at").defaultNow().notNull(),
})

export const workloadTag = pgTable("workload_tag", {
  workloadId: uuid()
    .notNull()
    .references(() => workloads.id, { onDelete: "cascade" }),
  tagId: bigint({ mode: "bigint" })
    .notNull()
    .references(() => tags.id, { onDelete: "cascade" }),
})

export const workloadRelations = relations(workloads, ({ many }) => ({
  environments: many(workloadEnvironment),
}))

export const workloadEnvironmentRelations = relations(workloadEnvironment, ({ one }) => ({
  workload: one(workloads, {
    fields: [workloadEnvironment.workloadId],
    references: [workloads.id],
  }),
  environment: one(environments, {
    fields: [workloadEnvironment.environmentId],
    references: [environments.id],
  }),
}))

export const workloadLensRelations = relations(workloadLens, ({ one }) => ({
  workload: one(workloads, {
    fields: [workloadLens.workloadId],
    references: [workloads.id],
  }),
  lens: one(lenses, {
    fields: [workloadLens.lensId],
    references: [lenses.id],
  }),
}))

export const workloadProfileRelations = relations(workloadProfile, ({ one }) => ({
  workload: one(workloads, {
    fields: [workloadProfile.workloadId],
    references: [workloads.id],
  }),
  profile: one(profiles, {
    fields: [workloadProfile.profileId],
    references: [profiles.id],
  }),
}))

export const workloadTagRelations = relations(workloadTag, ({ one }) => ({
  workload: one(workloads, {
    fields: [workloadTag.workloadId],
    references: [workloads.id],
  }),
  tag: one(tags, {
    fields: [workloadTag.tagId],
    references: [tags.id],
  }),
}))

export const workloadInsertSchema = createInsertSchema(workloads, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
  description: (schema) =>
    schema.min(1, "Description is required").max(1024, "Description must be at most 1024 characters"),
}).pick({
  name: true,
  description: true,
})

export const workloadSelectSchema = createSelectSchema(workloads)
export const workloadDeleteSchema = createSelectSchema(workloads).pick({
  id: true,
})

export type TWorkloadInsertSchema = ReturnType<typeof workloadInsertSchema.parse>
export type TWorkloadSelectSchema = ReturnType<typeof workloadSelectSchema.parse>
export type TWorkloadDeleteSchema = ReturnType<typeof workloadDeleteSchema.parse>
