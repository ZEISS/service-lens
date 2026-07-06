import { getDesigns } from "@/db/queries/designs"
import { paginationParams } from "@/db/queries/pagination"
import type { SearchParams } from "@/types"

import { DesignDataTable } from "./_components/data-table"

interface IndexPageProps {
  searchParams: Promise<SearchParams>
}

export const revalidate = 0

export default async function Page({ searchParams }: IndexPageProps) {
  const params = await searchParams
  const parsedParams = paginationParams.parse(params)

  const promises = Promise.all([
    getDesigns({
      ...parsedParams,
    }),
  ])

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      <DesignDataTable promises={promises} />
    </div>
  )
}
