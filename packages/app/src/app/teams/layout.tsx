import type { PropsWithChildren } from 'react'
import { cookies } from 'next/headers'
import { MainNav } from '@/components/teams/main-nav'
import DefaultLayout from '@/components/default-layout'

export default function Layout({ children }: PropsWithChildren) {
  const cookiesList = cookies()
  const hasScope = cookiesList.has('scope')
  const scope = cookiesList.get('scope')

  return (
    <>
      <DefaultLayout
        fallback={<MainNav scope={scope?.value} className="mx-6" />}
      >
        {children}
      </DefaultLayout>
    </>
  )
}
