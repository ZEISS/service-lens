import { PropsWithChildren } from 'react'
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

export default async function Page(props: PropsWithChildren<NextPageProps>) {
  return <>Dashboard</>
}
