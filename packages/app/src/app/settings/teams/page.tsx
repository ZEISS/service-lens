import { Separator } from '@/components/ui/separator'
import { columns } from '@/components/settings/teams-data-columns'
import { type DataTableOptions } from '@/components/data-table'
import { TeamsListSchema } from '@/server/routers/schemas/teams'
import { api } from '@/trpc/server-http'
import { PropsWithChildren } from 'react'
import { DataTable } from '@/components/data-table'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export interface NextPageProps {
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page(props: NextPageProps) {
  const searchParams = TeamsListSchema.parse({
    ...props.searchParams
  })
  const { rows, count } = await api.teams.list.query(searchParams)
  const pageCount = Math.ceil(count / searchParams.limit)

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Teams</h3>
        <p className="text-sm text-muted-foreground"></p>
      </div>
      <Separator />
      <DataTable
        columns={columns}
        data={rows}
        pageCount={pageCount}
        pageSize={searchParams.limit}
        pageIndex={searchParams.offset}
        options={options}
      />
    </div>
  )
}
