import { AddProfileButton } from '@/components/profiles/add-profile'
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
import { columns } from '@/components/dashboard/profiles/data-columns'
import type { DataTableOptions } from '@/components/data-table'
import { ListProfileByTeamSlug } from '@/server/routers/schemas/profile'

const options = {
  toolbar: {}
} satisfies DataTableOptions

export const revalidate = 0 // no cache

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default async function Page(props: PropsWithChildren<NextPageProps>) {
  const searchParams = ListProfileByTeamSlug.parse({
    ...props.searchParams,
    slug: props.params.team
  })
  const { rows, count } = await api.profiles.listByTeam.query(searchParams)
  const pageCount = Math.ceil(count / searchParams.limit)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          Profiles
          <SubNavSubtitle>Add business context to workloads</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions>
          <AddProfileButton teamSlug={props.params.team} />
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
