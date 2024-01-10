'use client'

import { columns } from './data-columns'
import { DataTable } from '@/components/data-table'
import { useQuery } from '@/lib/api'
import { api } from '@/trpc/client'

export function ProfilesDataTable() {
  const query = useQuery(({ pageIndex: offset, pageSize: limit }) =>
    api.profiles.list.query({ offset, limit })
  )

  return (
    <div className="h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <DataTable columns={columns} query={query()} />
    </div>
  )
}
