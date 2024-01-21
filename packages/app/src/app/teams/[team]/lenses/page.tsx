import { AddLensButton } from '@/components/lenses/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { Main } from '@/components/main'
import { api } from '@/trpc/server-http'
import type { PropsWithChildren } from 'react'
import { DataTable } from '@/components/data-table'
import { columns } from '@/components/lenses/data-columns'
import type { DataTableOptions } from '@/components/data-table'
import { ListLensByTeamSlug } from '@/server/routers/schemas/lens'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export const revalidate = 0 // no cache

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page(props: PropsWithChildren<NextPageProps>) {
  const searchParams = ListLensByTeamSlug.parse({
    ...props.searchParams,
    slug: props.params.team
  })
  const { rows, count } = await api.lenses.listByTeam.query(searchParams)

  const pageCount = Math.ceil(count / searchParams.limit)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          Lenses
          <SubNavSubtitle>
            Measure any architecture against best practices.
          </SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddLensButton />
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
