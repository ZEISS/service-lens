import { AddSolutionButton } from '@/components/solutions/add-solution'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { PropsWithChildren } from 'react'
import { Main } from '@/components/main'
import { api } from '@/trpc/server-http'
import { DataTable } from '@/components/data-table'
import { columns } from '@/components/solutions/data-columns'
import type { DataTableOptions } from '@/components/data-table'
import { ListSolutionByTeamSlug } from '@/server/routers/schemas/solution'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export const revalidate = 0 // no cache

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page(props: PropsWithChildren<NextPageProps>) {
  const searchParams = ListSolutionByTeamSlug.parse({
    ...props.searchParams,
    slug: props.params.team
  })
  const { rows, count } = await api.solutions.listByTeam.query(searchParams)
  const pageCount = Math.ceil(count / searchParams.limit)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          Solutions
          <SubNavSubtitle>Manage and discuss solutions</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddSolutionButton teamSlug={props.params.team} />
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
