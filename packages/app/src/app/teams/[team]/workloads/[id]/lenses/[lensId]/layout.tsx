import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { SidebarNav } from '@/components/dashboard/lenses/sidebar-nav'
import React from 'react'
import { api } from '@/trpc/server-http'
import type { PropsWithChildren } from 'react'

export type LayoutProps = {
  params: { id: string; lensId: string }
}

export default async function Layout({
  params,
  children
}: PropsWithChildren<LayoutProps>) {
  const lens = await api.getLens.query(params?.lensId)
  const workload = await api.workloads.get.query(params?.id)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          {workload?.name}
          <SubNavSubtitle>{lens?.name}</SubNavSubtitle>
        </SubNavTitle>
        <SubNavActions></SubNavActions>
      </SubNav>
      <main className="p-8">
        <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
          <aside className="lg:w-1/5">
            {lens && <SidebarNav lens={lens} params={params} />}
          </aside>
          <div className="flex-1">
            <div className="space-y-6">{children}</div>
          </div>
        </div>
      </main>
    </>
  )
}
