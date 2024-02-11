import { AddWorkloadButton } from '@/components/workloads/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { api } from '@/trpc/server-http'
import { DataTable } from '@/components/data-table'
import { columns } from '@/components/workloads/columns'
import type { DataTableOptions } from '@/components/data-table'
import { ListWorkloadByTeamSlug } from '@/server/routers/schemas/workload'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export const revalidate = 0 // no cache
export const dynamic = 'force-dynamic'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page(props: NextPageProps) {
  const searchParams = ListWorkloadByTeamSlug.parse({
    ...props.searchParams,
    slug: props.params.team
  })
  const { rows, count } = await api.workloads.listByTeam.query(searchParams)
  const pageCount = Math.ceil(count / searchParams.limit)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          Workloads
          <SubNavSubtitle>Manage and review workflows</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddWorkloadButton teamSlug={props.params.team} />
        </SubNavActions>
      </SubNav>
      <Main className="space-y-8 p-8">
        <DataTable
          columns={columns}
          data={rows}
          pageCount={pageCount}
          pageSize={searchParams.limit}
          pageIndex={searchParams.offset}
          options={options}
        />
      </Main>
    </>
  )
}
