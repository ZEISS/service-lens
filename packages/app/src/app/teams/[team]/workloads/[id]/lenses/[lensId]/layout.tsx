import { SubNav, SubNavTitle, SubNavActions } from '@/components/sub-nav'
import { SidebarNav } from '@/components/lenses/sidebar-nav'
import React from 'react'
import { api } from '@/trpc/server-http'
import { LoadingSpinner } from '@/components/loading-spinner'
import { type PropsWithChildren, Suspense } from 'react'

export const dynamic = 'force-dynamic'

export interface NextPageProps<TeamSlug = string, WorkloadId = string> {
  params: { team: TeamSlug; id: WorkloadId; questionId: string; lensId: string }
}

export default async function Layout({
  params,
  children
}: PropsWithChildren<NextPageProps>) {
  const workload = await api.workloads.get.query(params?.id)

  return (
    <>
      <SubNav>
        <SubNavTitle>{workload?.name}</SubNavTitle>
        <SubNavActions></SubNavActions>
      </SubNav>
      <main className="p-8">
        <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
          <aside className="lg:w-1/5">
            <Suspense fallback={<LoadingSpinner />}>
              <SidebarNav params={params} />
            </Suspense>
          </aside>
          <div className="flex-1">
            <div className="space-y-6">{children}</div>
          </div>
        </div>
      </main>
    </>
  )
}
