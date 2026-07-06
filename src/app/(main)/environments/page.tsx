import { getEnvironments } from "@/db/queries/environments"
import { paginationParams } from "@/db/queries/pagination"
import type { SearchParams } from "@/types"

import { EnvironmentDataTable } from "./_components/data-table"

interface IndexPageProps {
  searchParams: Promise<SearchParams>
}

export const revalidate = 0

export default async function Page({ searchParams }: IndexPageProps) {
  const params = await searchParams
  const parsedParams = paginationParams.parse(params)

  const promises = Promise.all([
    getEnvironments({
      ...parsedParams,
    }),
  ])

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      <EnvironmentDataTable promises={promises} />
    </div>
  )
}
