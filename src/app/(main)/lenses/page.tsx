import { getLenses } from "@/db/queries/lenses"
import { paginationParams } from "@/db/queries/pagination"
import type { SearchParams } from "@/types"
import { LensDataTable } from "./_components/data-table"

interface IndexPageProps {
  searchParams: Promise<SearchParams>
}

export const revalidate = 0

export default async function Page({ searchParams }: IndexPageProps) {
  const params = await searchParams
  const parsedParams = paginationParams.parse(params)

  const promises = Promise.all([
    getLenses({
      ...parsedParams,
    }),
  ])

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      <LensDataTable promises={promises} />
    </div>
  )
}
