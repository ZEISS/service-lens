import { SidebarNav } from '@/components/sidebar-nav'
import { PropsWithChildren } from 'react'

export type LayoutProps = {
  children?: React.ReactNode
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

export default function Layout({ children }: PropsWithChildren<LayoutProps>) {
  return (
    <>
      <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
        <SidebarNav items={sidebarNavItems} />
        <div className="flex-1 lg:max-w-2xl">
          <div className="space-y-6">{children}</div>
        </div>
      </div>
    </>
  )
}
