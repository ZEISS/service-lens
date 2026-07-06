import { parseAsInteger } from "nuqs/server"
import * as z from "zod"

export const paginationParams = z.object({
  page: z.coerce.number().min(1).default(1),
  perPage: z.coerce.number().min(1).default(10),
  order: z.enum(["asc", "desc"]).default("asc"),
  orderBy: z.string().default("id"),
})

export const searchParamsSchema = z.object({
  page: parseAsInteger.withDefault(1),
  perPage: parseAsInteger.withDefault(10),
  order: z.enum(["asc", "desc"]).default("asc"),
  orderBy: z.string().default("id"),
})

export type PaginationSchema = z.infer<typeof paginationParams>
