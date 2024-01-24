import { Button } from '@/components/ui/button'
import Link from 'next/link'
import { PropsWithChildren } from 'react'

export interface AddSolutionButtonProps {
  teamSlug: string
}

export function AddSolutionButton({
  teamSlug
}: PropsWithChildren<AddSolutionButtonProps>) {
  return (
    <Link href={`/teams/${teamSlug}/solutions/new`} passHref>
      <Button variant={'outline'}>Add Solution</Button>
    </Link>
  )
}
