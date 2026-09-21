import { headers } from "next/headers"

import { paginationParams } from "@/db/queries/pagination"
import { getWorkloads } from "@/db/queries/workloads"
import type { TWorkload } from "@/db/schema"
import type { TDesign } from "@/db/schemas/design"
import { auth } from "@/lib/auth"

export type CreateDesignRequest = Omit<TDesign, "id" | "createdAt" | "updatedAt">
export type UpdateDesignRequest = Partial<Omit<TDesign, "id">> & { id: TDesign["id"] }
export type DeleteDesignRequest = { id: TDesign["id"] }

export type ApiRequest<T = void> = {
  offset?: number
  limit?: number
  search?: string
  filter?: Partial<T>
}

export type ApiResponse<T> = {
  items: T[]
  totalCount: number
  offset: number
  limit: number
}
export const ApiResponse = class<T> implements ApiResponse<T> {
  items: T[] = []
  totalCount = 0
  offset = 0
  limit = 0
}

export type GetWorkloadsResponse = ApiResponse<TWorkload>
export type GetWorkloadsRequest = ApiRequest<TWorkload>

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)

  const session = await auth.api.getSession({
    headers: await headers(),
  })

  if (!session) {
    return new Response(JSON.stringify({ error: "User is not authenticated" }), {
      status: 401,
      headers: { "Content-Type": "application/json" },
    })
  }

  const parsedParams = paginationParams.parse(Object.fromEntries(searchParams))
  const workloads = await getWorkloads(parsedParams)

  const res = new ApiResponse<TWorkload>()
  res.items = workloads.data
  res.totalCount = workloads.data.length
  res.offset = parsedParams.page
  res.limit = parsedParams.perPage

  return new Response(JSON.stringify(res), {
    status: 200,
    headers: { "Content-Type": "application/json" },
  })
}
