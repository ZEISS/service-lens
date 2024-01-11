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
      <DefaultLayout>{children}</DefaultLayout>
    </>
  )
}
