'use client'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { use } from 'react'
import { api } from '@/trpc/client'
import { MagicWandIcon } from '@radix-ui/react-icons'
import { WorkloadDataTable } from '@/components/workloads/data-table'

export default function WorkloadsListCard() {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">Workloads</CardTitle>
      </CardHeader>
      <CardContent className="py-4">
        <WorkloadDataTable />
      </CardContent>
    </Card>
  )
}
