import {
  SubNav,
  SubNavTitle,
  SubNavActions,
  SubNavSubtitle
} from '@/components/sub-nav'
import { SidebarNav } from '@/components/sidebar-nav'
import { Main } from '@/components/main'
import DefaultLayout from '@/components/default-layout'

type PageProps = {
  children?: React.ReactNode
}

const sidebarNavItems = [
  {
    title: 'Profile',
    href: '/account'
  },
  {
    title: 'Appearance',
    href: '/account/appearance'
  }
]

export default function Layout({ children }: PageProps) {
  return (
    <>
      <DefaultLayout>
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
      </DefaultLayout>
    </>
  )
}
