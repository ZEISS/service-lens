import { paginationParams } from "@/db/queries/pagination"
import { getWorkloads } from "@/db/queries/workloads"
import type { SearchParams } from "@/types"

import { WorkloadDataTable } from "../../workloads/_components/data-table"
import { SectionCards } from "./_components/section-cards"

interface IndexPageProps {
  searchParams: Promise<SearchParams>
}

export const revalidate = 0

export default async function Page({ searchParams }: IndexPageProps) {
  const params = await searchParams
  const parsedParams = paginationParams.parse(params)

  const promises = Promise.all([
    getWorkloads({
      ...parsedParams,
    }),
  ])

  return (
    <div className="@container/main flex flex-col gap-4 md:gap-6">
      <SectionCards />
      <WorkloadDataTable promises={promises} />
    </div>
  )
}
