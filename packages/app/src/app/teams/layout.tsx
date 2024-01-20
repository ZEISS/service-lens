import type { PropsWithChildren } from 'react'
import { MainNav } from '@/components/teams/main-nav'
import DefaultLayout from '@/components/default-layout'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
  searchParams?: { [key: string]: string | string[] | undefined }
}

export default function Layout({
  params,
  children
}: PropsWithChildren<NextPageProps>) {
  return (
    <>
      <DefaultLayout
        fallback={<MainNav teamId={params.team} className="mx-6" />}
      >
        {children}
      </DefaultLayout>
    </>
  )
}
