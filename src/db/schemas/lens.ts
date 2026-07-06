import { pgTable } from "@/db/utils"
import { integer, json, timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

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
  version: (schema) => schema.int().min(1, "Version must be a positive integer"),
  description: (schema) => schema.max(1024, "Description must be at most 1024 characters").optional(),
}).pick({
  name: true,
  version: true,
  raw: true,
})

export const lensSelectSchema = createSelectSchema(lenses)
export const lensDeleteSchema = createSelectSchema(lenses).pick({
  id: true,
})

export type TLensInsertSchema = ReturnType<typeof lensInsertSchema.parse>
export type TLensSelectSchema = ReturnType<typeof lensSelectSchema.parse>
export type TLensDeleteSchema = ReturnType<typeof lensDeleteSchema.parse>
