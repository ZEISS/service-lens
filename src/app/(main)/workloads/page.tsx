import { paginationParams } from "@/db/queries/pagination"
import { listWorkloadsWithPagination } from "@/db/queries/workloads"
import type { SearchParams } from "@/types"

import { WorkloadDataTable } from "./_components/data-table"

interface IndexPageProps {
  searchParams: Promise<SearchParams>
}

export const revalidate = 0

export default async function Page({ searchParams }: IndexPageProps) {
  const params = await searchParams
  const parsedParams = paginationParams.parse(params)

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      <WorkloadDataTable promises={Promise.all([listWorkloadsWithPagination({...parsedParams})])} />
    </div>
  )
}
