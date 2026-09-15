import "dotenv/config"

import { defineConfig } from "drizzle-kit"

const databaseUrl = process.env.AZURE_POSTGRESQL_CONNECTIONSTRING ?? process.env.DATABASE_URL ?? ''

export default defineConfig({
  out: "./drizzle",
  schema: "./src/db/schema.ts",
  dialect: "postgresql",
  dbCredentials: {
    url: databaseUrl,
  },
})
