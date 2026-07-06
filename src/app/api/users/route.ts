import type { TDesign } from "@/db/schemas/design"

export type CreateDesignRequest = Omit<TDesign, "id" | "createdAt" | "updatedAt">
export type UpdateDesignRequest = Partial<Omit<TDesign, "id">> & { id: TDesign["id"] }
export type DeleteDesignRequest = { id: TDesign["id"] }

export type ApiRequest<T = void> = {
  offset?: number
  limit?: number
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

export type GetDesignsRequest = ApiRequest<TDesign>

export async function GET(request: Request) {
  const { searchParams } = new URL(request.url)
  const offset = Number(searchParams.get("offset")) || 0
  const limit = Number(searchParams.get("limit")) || 100

  // TODO: Replace with actual database query using getDesigns
  const designs: TDesign[] = []

  const res = new ApiResponse<TDesign>()
  res.items = designs
  res.totalCount = designs.length
  res.offset = offset
  res.limit = limit

  return new Response(JSON.stringify(res), {
    status: 200,
    headers: { "Content-Type": "application/json" },
  })
}

export async function POST(request: Request) {
  const body = (await request.json()) as CreateDesignRequest
  // TODO: Replace with actual database insertion using insertDesign
  const design: TDesign = {
    id: "temp-id",
    title: body.title,
    body: body.body,
    description: body.description,
    createdAt: new Date(),
    updatedAt: new Date(),
    deletedAt: null,
  }

  return new Response(JSON.stringify(design), {
    status: 201,
    headers: { "Content-Type": "application/json" },
  })
}
