'use client'

import { columns } from './data-columns'
import { DataTable, Query } from '@/components/data-table'
import type { Solution } from '@/db/models/solution'
import { api } from '@/trpc/client'

const query: Query<Solution> = ({ pageIndex: offset, pageSize: limit }) =>
  api.solutions.list.query({ offset, limit })

export default function SolutionsDataTable() {
  return (
    <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <DataTable columns={columns} query={query} />
    </div>
  )
}
