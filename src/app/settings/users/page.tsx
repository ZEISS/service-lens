import { api } from '@/trpc/server-http'
import { columns } from '@/components/settings/users-data-columns'
import { DataTable } from '@/components/data-table'
import { ListUsersSchema } from '@/server/routers/schemas/users'
import { Separator } from '@/components/ui/separator'
import { type DataTableOptions } from '@/components/data-table'

export const dynamic = 'force-dynamic'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export interface NextPageProps<TeamSlug = string> {
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page(props: NextPageProps) {
  const searchParams = ListUsersSchema.parse({
    ...props.searchParams
  })
  const { rows, count } = await api.users.list.query(searchParams)
  const pageCount = Math.ceil(count / searchParams.limit)

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-lg font-medium">Users</h3>
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
