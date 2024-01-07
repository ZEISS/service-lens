import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { SidebarNav } from './components/sidebar-nav'
import React from 'react'
import { api } from '@/trpc/server-http'

type PageProps = {
  children?: React.ReactNode
}

const sidebarNavItems = [
  {
    title: 'General',
    href: '/dashboard/settings'
  },
  {
    title: 'Environments',
    href: '/dashboard/settings/environments'
  },
  {
    title: 'Developer',
    href: '/dashboard/settings/developer'
  }
]

export type LayoutProps = {
  children?: React.ReactNode
  params: { id: string; lensId: string }
}

export default async function Layout({ params, children }: LayoutProps) {
  const lens = await api.getLens.query(params?.lensId)

  return (
    <>
      <SubNav>
        <SubNavTitle>
          {lens?.name}
          <SubNavSubtitle>{lens?.description}</SubNavSubtitle>
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
