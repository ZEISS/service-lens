import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { PropsWithChildren } from 'react'

export interface AddLensButton {
  teamSlug: string
}

export function AddLensButton({ teamSlug }: PropsWithChildren<AddLensButton>) {
  return (
    <Link href={`/teams/${teamSlug}/lenses/new`} passHref>
      <Button variant={'outline'}>Add Lens</Button>
    </Link>
  )
}
