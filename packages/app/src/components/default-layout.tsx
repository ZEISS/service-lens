import '@/styles/globals.css'
import type { ReactNode } from 'react'

import { MainNav } from '@/components/main-nav'
import { Search } from '@/components/search'
import TeamSwitcher from '@/components/team-switcher'
import { UserNav } from '@/components/user-nav'
import { ThemeToggle } from '@/components/theme-toggle'
import { Toaster } from '@/components/ui/toaster'
import Footer from '@/components/footer'

interface DefaultLayoutProps {
  children?: ReactNode | undefined

  /** A fallback react tree to show when a Suspense child (like React.lazy) suspends */
  fallback?: ReactNode
}

export default function DefaultLayout({
  children,
  fallback = <MainNav className="mx-6" />
}: DefaultLayoutProps) {
  return (
    <>
      <div className="flex-col">
        <div className="border-b">
          <div className="flex h-16 items-center px-4">
            <TeamSwitcher />
            {fallback}
            <div className="ml-auto flex items-center space-x-4">
              <Search />
              <ThemeToggle />
              <UserNav />
            </div>
          </div>
        </div>
        <div className="flex-1">{children}</div>
        <Footer />
      </div>
      <Toaster />
    </>
  )
}
