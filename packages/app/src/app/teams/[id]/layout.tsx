import { SidebarNav } from '@/components/sidebar-nav'
import { SubNav, SubNavTitle, SubNavActions } from '@/components/sub-nav'
import type { PropsWithChildren } from 'react'
import { api } from '@/trpc/server-invoker'

export type LayoutProps = {
  params: { id: string }
}

const sidebarNavItems = [
  {
    title: 'Members',
    href: '/account'
  },
  {
    title: 'Settings',
    href: '/account/appearance'
  },
  {
    title: 'Notifications',
    href: '/account/notifications'
  }
]

export default async function Layout({
  children,
  params
}: PropsWithChildren<LayoutProps>) {
  const team = await api.teams.get.query(params.id)

  return (
    <>
      <SubNav>
        <SubNavTitle>{team?.name}</SubNavTitle>
        <SubNavActions></SubNavActions>
      </SubNav>
      <main className="p-8">
        <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
          <aside className="-mx-4 lg:w-1/5">
            <SidebarNav items={sidebarNavItems} />
          </aside>
          <div className="flex-1 lg:max-w-2xl">
            <div className="space-y-6">{children}</div>
          </div>
        </div>
      </main>
    </>
  )
}
