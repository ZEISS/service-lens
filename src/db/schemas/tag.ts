import { pgTable } from "@/db/utils"
import { bigserial, index, timestamp, uniqueIndex, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"

export const tags = pgTable(
  "tag",
  {
    id: bigserial({ mode: "number" }).primaryKey(),
    name: varchar({ length: 255 }).notNull(),
    value: varchar({ length: 1024 }).notNull(),
    createdAt: timestamp("created_at").defaultNow(),
    updatedAt: timestamp("updated_at")
      .defaultNow()
      .$onUpdate(() => new Date()),
    deletedAt: timestamp("deleted_at"),
  },
  (table) => [
    index("tag_name_index").on(table.name),
    uniqueIndex("tag_name_value_unique_index").on(table.name, table.value), // Ensure unique combination of name and value
  ],
)

export type TTag = typeof tags.$inferSelect
export type TNewTag = typeof tags.$inferInsert

export const tagInsertSchema = createInsertSchema(tags, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
  value: (schema) => schema.min(1, "Value is required").max(1024, "Value must be at most 1024 characters"),
}).pick({
  name: true,
  value: true,
})
export const tagSelectSchema = createSelectSchema(tags)
export const tagDeleteSchema = createSelectSchema(tags).pick({
  id: true,
})

export type TTagInsertSchema = ReturnType<typeof tagInsertSchema.parse>
export type TTagSelectSchema = ReturnType<typeof tagSelectSchema.parse>
export type TTagDeleteSchema = ReturnType<typeof tagDeleteSchema.parse>
