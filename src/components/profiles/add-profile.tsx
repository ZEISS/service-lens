import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { PropsWithChildren } from 'react'

export interface AddProfileButton {
  teamSlug: string
}

export function AddProfileButton({
  teamSlug
}: PropsWithChildren<AddProfileButton>) {
  return (
    <Link href={`/teams/${teamSlug}/profiles/new`} passHref>
      <Button variant={'outline'}>Add Profile</Button>
    </Link>
  )
}
