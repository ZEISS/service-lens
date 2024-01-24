import { SubNavTitle, SubNav, SubNavActions } from '@/components/sub-nav'
import { SidebarNav } from '@/components/sidebar-nav'
import { Main } from '@/components/main'
import DefaultLayout from '@/components/default-layout'
import { PropsWithChildren } from 'react'

export interface NextPageProps {
  params: {}
  searchParams?: { [key: string]: string | string[] | undefined }
}

const sidebarNavItems = [
  {
    title: 'General',
    href: '/settings'
  },
  {
    title: 'Teams',
    href: '/settings/teams'
  }
]

export default function Layout({ children }: PropsWithChildren<NextPageProps>) {
  return (
    <>
      <DefaultLayout>
        <SubNav>
          <SubNavTitle>Settings</SubNavTitle>
          <SubNavActions></SubNavActions>
        </SubNav>
        <Main className="p-8">
          <div className="flex flex-col space-y-8 lg:flex-row lg:space-x-12 lg:space-y-0">
            <SidebarNav items={sidebarNavItems} />
            <div className="flex-1 lg:max-w-2xl">
              <div className="space-y-6">{children}</div>
            </div>
          </div>
        </Main>
      </DefaultLayout>
    </>
  )
}
