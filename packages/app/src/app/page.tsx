import { Metadata } from 'next'
import { auth } from '@/auth'
import Link from 'next/link'
import { redirect } from 'next/navigation'
import DefaultLayout from '@/components/default-layout'

export const metadata: Metadata = {
  title: 'Root',
  description: 'Root'
}

export default async function Root() {
  const session = await auth()

  if (session !== null) {
    redirect('/dashboard')
  }

  return (
    <>
      <DefaultLayout>
        <Link href="/login">Login</Link>
      </DefaultLayout>
    </>
  )
}
