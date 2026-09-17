import "@/config/env-config"

import "dotenv/config"
export const DATABASE_PREFIX = process.env.DATABASE_PREFIX ?? "service_lens"
