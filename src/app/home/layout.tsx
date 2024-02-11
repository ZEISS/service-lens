import DefaultLayout from '@/components/default-layout'
import React from 'react'

export type PageLayoutProps = {
  children?: React.ReactNode
}

export default function PageLayout({ children }: PageLayoutProps) {
  return (
    <>
      <DefaultLayout>{children}</DefaultLayout>
    </>
  )
}
