import { headers } from "next/headers"

import { paginationParams } from "@/db/queries/pagination"
import { getProfiles } from "@/db/queries/profiles"
import type { TProfile } from "@/db/schema"
import { auth } from "@/lib/auth"

export type CreateProfileRequest = Omit<TProfile, "id" | "createdAt" | "updatedAt">
export type UpdateProfileRequest = Partial<Omit<TProfile, "id">> & { id: TProfile["id"] }
export type DeleteProfileRequest = { id: TProfile["id"] }

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

export type GetProfileResponse = ApiResponse<TProfile>
export type GetProfileRequest = ApiRequest<TProfile>

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
  const profiles = await getProfiles(parsedParams)

  const res = new ApiResponse<TProfile>()
  res.items = profiles.data
  res.totalCount = profiles.data.length
  res.offset = parsedParams.page
  res.limit = parsedParams.perPage

  return new Response(JSON.stringify(res), {
    status: 200,
    headers: { "Content-Type": "application/json" },
  })
}
