'use client'

import { columns } from './columns'
import { DataTable } from '@/components/data-table'
import { useQuery } from '@/lib/api'
import { api } from '@/trpc/client'
import type { DataTableOptions } from '@/components/data-table'

const options = {
  toolbar: {
    facetFilters: [
      {
        column: 'environment',
        title: 'Environment',
        options: [
          { label: 'Active', value: 'active' },
          { label: 'Inactive', value: 'inactive' }
        ]
      }
    ]
  }
} satisfies DataTableOptions

export function WorkloadDataTable() {
  const query = useQuery(({ pageIndex: offset, pageSize: limit }) =>
    api.listWorkloads.query({ offset, limit })
  )

  return (
    <div className="hidden h-full flex-1 flex-col space-y-8 p-8 md:flex">
      <DataTable columns={columns} query={query()} options={options} />
    </div>
  )
}
