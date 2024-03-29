import '@/styles/globals.css'
import { Suspense, type ReactNode } from 'react'

import { cookies } from 'next/headers'
import { MainNav } from '@/components/main-nav'
import { Search } from '@/components/search'
import { ThemeToggle } from '@/components/theme-toggle'
import { Toaster } from '@/components/ui/toaster'
import { UserNav } from '@/components/user-nav'
import Footer from '@/components/footer'
import TeamSwitcher from '@/components/team-switcher'
import { api } from '@/trpc/server-http'

interface DefaultLayoutProps {
  children?: ReactNode | undefined
  fallback?: ReactNode
}

export default async function DefaultLayout({
  children,
  fallback = <MainNav className="mx-6" />
}: DefaultLayoutProps) {
  const cookiesList = cookies()
  const scope = cookiesList.get('scope')
  const user = await api.users.get.query()

  return (
    <>
      <div className="flex-col">
        <div className="border-b">
          <div className="flex h-16 items-center px-4">
            <TeamSwitcher scope={scope?.value} user={user} />
            {fallback}
            <div className="ml-auto flex items-center space-x-4">
              <Search />
              <ThemeToggle />
              <UserNav user={user} />
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
