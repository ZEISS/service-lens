import { AddWorkloadButton } from '@/components/workloads/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { PropsWithChildren } from 'react'
import { Main } from '@/components/main'
import { z } from 'zod'
import { api } from '@/trpc/server-http'
import { DataTable } from '@/components/data-table'
import { columns } from '@/components/workloads/columns'
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

export const revalidate = 0 // no cache

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

const searchParamsSchema = z.object({
  per_page: z.coerce.number().default(10),
  page: z.coerce.number().default(0)
})

export default async function Page(props: PropsWithChildren<NextPageProps>) {
  const searchParams = searchParamsSchema.parse(props.searchParams)
  const { rows, count } = await api.workloads.listByTeam.query({
    slug: props.params.team,
    ...searchParams
  })

  const pageCount = Math.ceil(count / searchParams.per_page)

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
          pageSize={searchParams.per_page}
          pageIndex={searchParams.page}
          options={options}
        />
      </Main>
    </>
  )
}
