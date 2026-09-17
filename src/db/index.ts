import "dotenv/config"
import { drizzle } from "drizzle-orm/node-postgres"
import { Pool } from "pg"

import { relations } from "./relations"

const databaseUrl = process.env.POSTGRES_URL ?? process.env.DATABASE_URL ?? ""

const client = new Pool({
  connectionString: databaseUrl,
})

export const db = drizzle({
  client,
  relations,
})
