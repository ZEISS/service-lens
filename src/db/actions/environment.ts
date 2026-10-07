"use server"

import { z } from "zod"

import { db } from "@/db"
import { environmentInsertSchema, environments, type TNewEnvironment } from "@/db/schema"

export async function createEnvironment(environment: TNewEnvironment) {
  try {
    const parsedEnvironment = environmentInsertSchema.parse(environment)
    await db.insert(environments).values(parsedEnvironment)
  } catch (err) {
    if (err instanceof z.ZodError) {
      return err.issues
    }
  }
}
