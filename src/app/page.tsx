import { Metadata } from 'next'
import Link from 'next/link'
import DefaultLayout from '@/components/default-layout'

export const metadata: Metadata = {
  title: 'Root',
  description: 'Root'
}

export default async function Root() {
  return (
    <>
      <DefaultLayout>
        <Link href="/login">Login</Link>
      </DefaultLayout>
    </>
  )
}
