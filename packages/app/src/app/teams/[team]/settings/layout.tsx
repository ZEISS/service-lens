import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { PropsWithChildren } from 'react'
import { SidebarNav } from '@/components/sidebar-nav'
import { Main } from '@/components/main'

export interface NextPageProps<TeamSlug = string> {
  params: { team: TeamSlug }
}

export default function Layout({
  children,
  params
}: PropsWithChildren<NextPageProps>) {
  console.log(params)
  const sidebarNavItems = [
    {
      title: 'General',
      href: `/teams/${params.team}/settings`
    },
    {
      title: 'Members',
      href: `/teams/${params.team}/settings/members`
    }
  ]

  return (
    <>
      <SubNav>
        <SubNavTitle>
          Settings
          <SubNavSubtitle>
            Manage the settings of the service lens.
          </SubNavSubtitle>
        </SubNavTitle>
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
    </>
  )
}
