'use client'

import { columns } from './columns'
import { DataTable } from '@/components/data-table'
import { useQuery } from '@/lib/api'
import { api } from '@/trpc/client'
import { Workload } from '@/db/models/workload'
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

export function WorkloadDataTable({ teamSlug }: { teamSlug: string }) {
  const query = useQuery<Workload>(({ pageIndex: offset, pageSize: limit }) =>
    api.teams.listWorkloads.query(teamSlug).then(team => ({
      count: team?.workloads.length ?? 0,
      rows: team?.workloads ?? []
    }))
  )

  return (
    <div className="h-full flex-1 flex-col md:flex">
      <DataTable columns={columns} query={query()} options={options} />
    </div>
  )
}
