/**
 * @see https://gist.github.com/rphlmr/0d1722a794ed5a16da0fdf6652902b15
 */

import { DATABASE_PREFIX } from "@/config/db-config"
import { pgTableCreator } from "drizzle-orm/pg-core"

/**
 * Allows a single database instance for multiple projects.
 * @see https://orm.drizzle.team/docs/goodies#multi-project-schema
 */
export const pgTable = pgTableCreator((name) => `${DATABASE_PREFIX}_${name}`)

export function takeFirstOrNull<TData>(data: TData[]) {
  return data[0] ?? null
}

export function takeFirstOrThrow<TData>(data: TData[], errorMessage?: string) {
  const first = takeFirstOrNull(data)

  if (!first) {
    throw new Error(errorMessage ?? "Item not found")
  }

  return first
}
