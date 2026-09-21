import { getLenses } from "@/db/queries/lenses"
import { paginationParams } from "@/db/queries/pagination"
import type { TEnvironment, TLens } from "@/db/schema"

export type CreateLensRequest = Omit<TLens, "id" | "createdAt" | "updatedAt">
export type UpdateLensRequest = Partial<Omit<TLens, "id">> & { id: TLens["id"] }
export type DeleteLensRequest = { id: TLens["id"] }

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

export type GetLensResponse = ApiResponse<TLens>
export type GetLensRequest = ApiRequest<TLens>

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)

  const parsedParams = paginationParams.parse(Object.fromEntries(searchParams))
  const lenses = await getLenses(parsedParams)

  const res = new ApiResponse<TEnvironment>()
  res.items = lenses.data
  res.totalCount = lenses.data.length
  res.offset = parsedParams.page
  res.limit = parsedParams.perPage

  return new Response(JSON.stringify(res), {
    status: 200,
    headers: { "Content-Type": "application/json" },
  })
}
