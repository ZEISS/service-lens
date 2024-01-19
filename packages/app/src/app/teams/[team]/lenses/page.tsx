import { AddLensButton } from '@/components/lenses/add-button'
import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { z } from 'zod'
import { Main } from '@/components/main'
import { api } from '@/trpc/server-http'
import type { PropsWithChildren } from 'react'
import { DataTable } from '@/components/data-table'
import { columns } from '@/components/lenses/data-columns'
import type { DataTableOptions } from '@/components/data-table'

const options = {
  toolbar: {}
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
  const { rows, count } = await api.lenses.listByTeam.query({
    slug: props.params.team,
    ...searchParams
  })

  const pageCount = Math.ceil(count / searchParams.per_page)

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
          pageSize={searchParams.per_page}
          pageIndex={searchParams.page}
          options={options}
        />
      </Main>
    </>
  )
}
