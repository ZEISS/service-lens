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

export default function Layout({ children }: PageProps) {
  return (
    <>
      <DefaultLayout>
        <SubNav>
          <SubNavTitle>
            Teams
            <SubNavSubtitle>Manage teams and team members.</SubNavSubtitle>
          </SubNavTitle>
          <SubNavActions></SubNavActions>
        </SubNav>
        <Main className="p-8">{children}</Main>
      </DefaultLayout>
    </>
  )
}
