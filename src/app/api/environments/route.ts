import { headers } from "next/headers"

import { getEnvironments } from "@/db/queries/environments"
import { paginationParams } from "@/db/queries/pagination"
import type { TEnvironment } from "@/db/schema"
import { auth } from "@/lib/auth"

export type CreateEnvironmentRequest = Omit<TEnvironment, "id" | "createdAt" | "updatedAt">
export type UpdateEnvironmentRequest = Partial<Omit<TEnvironment, "id">> & { id: TEnvironment["id"] }
export type DeleteEnvironmentRequest = { id: TEnvironment["id"] }

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

export type GetEnvironmentResponse = ApiResponse<TEnvironment>
export type GetEnvironmentRequest = ApiRequest<TEnvironment>

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
  const environments = await getEnvironments(parsedParams)

  const res = new ApiResponse<TEnvironment>()
  res.items = environments.data
  res.totalCount = environments.data.length
  res.offset = parsedParams.page
  res.limit = parsedParams.perPage

  return new Response(JSON.stringify(res), {
    status: 200,
    headers: { "Content-Type": "application/json" },
  })
}
