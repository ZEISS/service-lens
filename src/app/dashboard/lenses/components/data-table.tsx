'use client'

import { columns } from './data-columns'
import { DataTable } from '@/components/data-table'
import type { Lens } from '@/db/models/lens'
import { useQuery } from '@/lib/api'
import { api } from '@/trpc/client'

export function LensesDataTable() {
  const query = useQuery(({ pageIndex: offset, pageSize: limit }) =>
    api.lenses.list.query({ offset, limit })
  )

  return (
    <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <DataTable<Lens> columns={columns} query={query()} />
    </div>
  )
}
