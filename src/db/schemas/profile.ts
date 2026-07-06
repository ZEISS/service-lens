import { pgTable } from "@/db/utils"
import { relations } from "drizzle-orm"
import { bigint, timestamp, uuid, varchar } from "drizzle-orm/pg-core"
import { createInsertSchema, createSelectSchema } from "drizzle-zod"
import { lenses } from "./lens"

export const profiles = pgTable("profile", {
  id: uuid().primaryKey().defaultRandom(),
  name: varchar({ length: 255 }).notNull(),
  description: varchar({ length: 1024 }),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TProfile = typeof profiles.$inferSelect
export type TNewProfile = typeof profiles.$inferInsert

export const profileLens = pgTable("profile_lens", {
  profileId: uuid()
    .notNull()
    .references(() => profiles.id, { onDelete: "cascade" }),
  lensId: uuid()
    .notNull()
    .references(() => lenses.id),
})

export const profileLensRelations = relations(profileLens, ({ many, one }) => ({
  profile: one(profiles, {
    fields: [profileLens.profileId],
    references: [profiles.id],
  }),
  lens: many(lenses),
}))

export const profileQuestion = pgTable("profile_question", {
  id: bigint({ mode: "bigint" }).primaryKey(),
  question: varchar({ length: 1024 }).notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TProfileQuestion = typeof profileQuestion.$inferSelect
export type TNewProfileQuestion = typeof profileQuestion.$inferInsert

export const profileQuestionAnswer = pgTable("profile_question_answer", {
  id: bigint({ mode: "bigint" }).primaryKey(),
  name: varchar({ length: 255 }).notNull(),
  profileQuestionId: bigint({ mode: "bigint" })
    .notNull()
    .references(() => profileQuestion.id),
  answer: varchar({ length: 2048 }).notNull(),
  createdAt: timestamp("created_at").defaultNow(),
  updatedAt: timestamp("updated_at")
    .defaultNow()
    .$onUpdate(() => new Date()),
  deletedAt: timestamp("deleted_at"),
})

export type TProfileQuestionAnswer = typeof profileQuestionAnswer.$inferSelect
export type TNewProfileQuestionAnswer = typeof profileQuestionAnswer.$inferInsert

export const profileRelations = relations(profiles, ({ many }) => ({
  questions: many(profileQuestion),
}))

export const profileQuestionRelations = relations(profileQuestion, ({ many, one }) => ({
  profile: one(profiles, {
    fields: [profileQuestion.id],
    references: [profiles.id],
  }),
  answers: many(profileQuestionAnswer),
}))

export const profileQuestionAnswerRelations = relations(profileQuestionAnswer, ({ one }) => ({
  question: one(profileQuestion, {
    fields: [profileQuestionAnswer.profileQuestionId],
    references: [profileQuestion.id],
  }),
}))

export const profileInsertSchema = createInsertSchema(profiles, {
  name: (schema) => schema.min(1, "Name is required").max(255, "Name must be at most 255 characters"),
}).pick({
  name: true,
})

export const profileSelectSchema = createSelectSchema(profiles)
export const profileDeleteSchema = createSelectSchema(profiles).pick({
  id: true,
})

export type TProfileInsertSchema = ReturnType<typeof profileInsertSchema.parse>
export type TProfileSelectSchema = ReturnType<typeof profileSelectSchema.parse>
export type TProfileDeleteSchema = ReturnType<typeof profileDeleteSchema.parse>
